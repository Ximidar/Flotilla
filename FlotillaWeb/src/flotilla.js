const axios = require("axios")
import { FileStructures } from './js_proto/FileStructures.js'

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
}

