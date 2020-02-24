const axios = require("axios")
import { FileStructures } from './js_proto/FileStructures_pb.js'
import { PlayStructures } from './js_proto/action_pb.js'
import { BufferReader } from 'protobufjs'

export class Flotilla{
    constructor(){
        this.base = "0.0.0.0:5000"
    }

    // Files
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

    async SelectFile(file){
        if (file.FileType != "file") {
            console.log("Not a file we can play")
            return
        }

        // turn file into object
        let proto_file = new FileStructures.File(file)
        console.log("proto")
        console.log(proto_file)
        let buf = new Uint8Array()
        buf = FileStructures.File.encode(proto_file).finish()
        console.log(buf)
        // post file action
        let url = "http://" + this.base + "/api/selectfile"
        axios.request({ responseType: 'text',
                        url: url,
                        data: Buffer.from(buf),
                        headers: { "content-type": "application/octet-stream",
                                    "blob-length": buf.length },
                        method: 'post'
        }).then(response => {
            console.log(response)
            return response
          })
          .catch(err => {
            console.log(err)
            return err
          })

        // we will have played a file!
    }


    // Status
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
        let act = new PlayStructures.Action(action_payload)
        let buf = new Uint8Array()
        buf = PlayStructures.Action.encode(act).finish()
        console.log(act)
        console.log(buf)
        let url = "http://" + this.base + "/api/status"
        axios.request({ responseType: 'text',
                        url: url,
                        data: Buffer.from(buf),
                        headers: { "content-type": "application/octet-stream",
                                    "blob-length": buf.length},
                        method: 'post'
        }).then(response => {
            console.log(response)
            return response
          })
          .catch(err => {
            console.log(err)
            return err
          })
        
    }

    
}

