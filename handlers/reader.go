// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/stack360/faas-lambdroid/lambdroid"
    "github.com/openfaas/faas/gateway/requests"
)

func MakeFunctionReader(towerClient lambdroid.LambdroidTowerClient) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        functions, err := getFunctions(towerClient)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        functionBytes, marshalErr := json.Marshal(functions)
        if marshalErr != nil {
            http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(functionBytes)
    }
}

func getFunctions(towerClient lambdroid.LambdroidTowerClient) ([]requests.Function, error) {
    functions := []requests.Function{}

    services, err := towerClient.ListServices()
    if err != nil {
        return nil, err
    }

    for _, service := range services {
        function := requests.Function{
            Name:              service,
            Replicas:          uint64(1),
            Image:             service,
            InvocationCount:   0,
            AvailableReplicas: 1,
        }
        functions = append(functions, function)
    }
    return functions, nil
}
