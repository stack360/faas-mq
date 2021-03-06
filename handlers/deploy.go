// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "encoding/json"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/openfaas/faas/gateway/requests"
    "github.com/stack360/faas-lambdroid/lambdroid"
)

func MakeDeployHandler(towerClient lambdroid.LambdroidTowerClient) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        log.Println("aaaaaaaa")
        defer r.Body.Close()
        body, _ := ioutil.ReadAll(r.Body)

        request := requests.CreateFunctionRequest{}
        err := json.Unmarshal(body, &request)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        serviceSpec := map[string]interface{} {}
        _, addErr := towerClient.AddService(serviceSpec)
        if addErr != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(addErr.Error()))
            return
        }
        w.WriteHeader(http.StatusAccepted)
    }
}
