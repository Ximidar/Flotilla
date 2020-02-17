const axios = require("axios")
import { FileStructures } from './js_proto/FileStructures.js'
import { Action } from './js_proto/action_pb.js'

export class Flotilla{
    constructor(){
        this.base = "0.0.0.0:5000"
    }

    async GetFiles(){
        var url = "http://" + this.base + "/api/getfiles"
        var req = axios.request({ responseType: 'blob',
                                  url: url,
                                  method: 'get'
            })
        var ab = await req
        var buf = await ab.data.arrayBuffer()
        let transbuf = new Uint8Array(buf)
        let file = FileStructures.File.decode(transbuf)
        console.log(file)
        return file
    }

    async GetStatus(){
        var url = "http://" + this.base + "/api/status"
        var req = axios.request({ responseType: 'text',
                                  url: url,
                                  method: 'get'
            })
        var ab = await req
        var data = await ab.data
        return data

    }

    async PostAction(action){
        var act = new Action.Action()
        act.setAction(action)
        let byteAction = act.serializeBinary()
        var url = "http://" + this.base + "/api/status"
        var req = axios.request({ responseType: 'text',
                                  url: url,
                                  data: byteAction,
                                  method: 'post'
        })
        var ab = await req
        var data = await ab.data
        console.log(data)
    }
}

