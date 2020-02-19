const axios = require("axios")
import { FileStructures } from './js_proto/FileStructures_pb.js'
import { PlayStructures } from './js_proto/action_pb.js'

export class Flotilla{
    constructor(){
        this.base = "0.0.0.0:5000"
    }

    async GetFiles(){
        let url = "http://" + this.base + "/api/getfiles"
        let req = axios.request({ responseType: 'blob',
                                  url: url,
                                  method: 'get'
            })
        let ab = await req
        let buf = await ab.data.arrayBuffer()
        let transbuf = new Uint8Array(buf)
        let file = FileStructures.File.decode(transbuf)
        console.log(file)
        return file
    }

    async GetStatus(){
        let url = "http://" + this.base + "/api/status"
        let req = axios.request({ responseType: 'text',
                                  url: url,
                                  method: 'get'
            })
        let ab = await req
        let data = await ab.data
        return data

    }

    async PostAction(action){
        let action_payload = {Action: action}
        let act = PlayStructures.Action.create(action_payload)
        let buffer = PlayStructures.Action.encode(act).finish()
        let url = "http://" + this.base + "/api/status"
        axios.request({ responseType: 'text',
                        url: url,
                        data: buffer,
                        headers: { "content-type": buffer.type,
                                    "blob-length": buffer.length - 1 },
                        method: 'post'
        }).then(response => {
            console.log(response)
            return response
          })
          .catch(err => {
            console.log(err)
          })
        
    }
}

