const axios = require("axios")
import { FileStructures } from './js_proto/FileStructures_pb.js'
import { PlayStructures } from './js_proto/action_pb.js'
import { CommStructures } from './js_proto/CommStructures_pb.js'
import { BufferReader } from 'protobufjs'

//const base = "0.0.0.0:5000"

export default {
    data(){
        return {
            base: "0.0.0.0:5000"
        }
    },
    methods: {
        // Status
        flotGetStatus: async function(){
            let url = "http://" + this.base + "/api/status"
            let req = axios.request({ responseType: 'text',
                                    url: url,
                                    method: 'get'
                })
            let ab = await req
            let data = await ab.data
            return data
        },

        flotPostAction: async function(action){
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
        },

        // Files
        flotGetFiles: async function(){
            let url = "http://" + this.base + "/api/getfiles"
            let req = axios.request({ responseType: 'blob',
                                    url: url,
                                    method: 'get'
                })
            let ab = await req
            let buf = await ab.data.arrayBuffer()
            let transbuf = new Uint8Array(buf)
            let file = FileStructures.File.decode(transbuf)
            return file
        },
        flotSelectFile: async function(file){
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
        },

        // Comm
        flotGetCommOptions: async function(){
            let url = "http://" + this.base + "/api/comm/options"
            let req = axios.request({
                responseType: 'blob',
                url: url,
                method: 'get'
            })

            let ab = await req
            let buf = await ab.data.arrayBuffer()
            let transbuf = new Uint8Array(buf)
            let commOptions = await CommStructures.CommOptions.decode(transbuf)
            return commOptions

        },
        flotCreateCommInit: function(port, baud){
            let InitComm = CommStructures.InitComm.create({
                port: port,
                baud: baud
            })
            return InitComm
        },
        flotSendCommInit: async function(commInit){
            let url = "http://" + this.base + "/api/comm/init"
            let ci = CommStructures.InitComm.encode(commInit).finish()
            axios.request({ responseType: 'blob',
                            url: url,
                            data: Buffer.from(ci),
                            headers: { "content-type": "application/octet-stream"},
                            method: 'post'
            }).then(response => {
                console.log(response)
                return response
            }).catch(err => {
                console.log(err)
                return err
            })

        },
        flotCommConnect: async function(){
            let url = "http://" + this.base + "/api/comm/connect"
            axios.request({
                responseType: 'blob',
                url: url,
                method: 'get'
            }).then(response =>{
                console.log(response)
                return response
            }).catch(err =>{
                console.log(err)
                return err
            })

        },
        flotCommDisconnect: async function(){
            let url = "http://" + this.base + "/api/comm/disconnect"
            axios.request({
                responseType: 'blob',
                url: url,
                method: 'get'
            }).then(response =>{
                console.log(response)
                return response
            }).catch(err =>{
                console.log(err)
                return err
            })
        }
    }
}