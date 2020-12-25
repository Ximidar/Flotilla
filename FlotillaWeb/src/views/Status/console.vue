<template>
    <v-container fill-height fill-width>
        <v-row align="center" justify="center">
            <v-col sm="6" cols="10">
                <v-toolbar><b>Console Monitor</b></v-toolbar>
                <v-sheet outlined >
                    <v-virtual-scroll
                        :items="console_text"
                        height="300"
                        item-height="24"
                    >
                        <template v-slot="{ item }">
                            <ConsoleItem :text="item"/>
                        </template>

                    </v-virtual-scroll>
                </v-sheet>
                <v-textarea
                    v-model="console_input"
                    :auto-grow="auto_clear_single"
                    :clearable="auto_clear_single"
                    :single-line="auto_clear_single"
                    filled
                    rows="1"
                    row-height="24"
                    @keyup="onEnter"

                ></v-textarea>
                
            </v-col>
            <v-col sm="6" cols="10">
                <PConsole/>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import ConsoleItem from "@/views/Status/console_item.vue"
import PConsole from "@/views/Status/printer_connection.vue"
export default {
    name: 'console',
    components: {
        ConsoleItem,
        PConsole,
    },
    data(){
        return{
            connected: false,
            console_text: ["line1", "line2", "line3"],
            console_history: [],
            history_pos: -1,
            temp_input: "",
            console_input: "",
            auto_clear_single: true,
            console_websocket: null
        }
    },
    methods: {
        onEnter: function(input){
            console.log("got input!")
            console.log(input)
            switch(input.key) {
                case "Enter": {
                    console.log("Got Enter")
                    console.log(this.console_input)
                    this.historyPush(this.console_input)
                    this.consolePush(this.console_input)
                    this.console_input = ""
                    this.history_pos = this.console_history.length

                }
                break;
                case "ArrowUp": {
                    console.log("arrow up")
                    this.history_pos--
                    if(this.history_pos < 0){
                        this.history_pos = -1
                        this.console_input=""
                        break
                    }
                    
                    this.console_input = this.console_history[this.history_pos]
                    
                    
                }
                break;
                case "ArrowDown": {
                    console.log("arrow down")
                    this.history_pos++
                    if (this.history_pos >= this.console_history.length){
                        this.history_pos = this.console_history.length
                        this.console_input=""
                        break
                    }
                    this.console_input = this.console_history[this.history_pos]

                }
                break;
            }
        },
        historyPush: function(item_input){
            this.console_history.push(item_input.replace(/[\s\n\r]/g, ''))
        },
        consolePush: function(item_input){
            // TODO push to flotilla console
            this.console_text.push(item_input.replace(/[\s\n\r]/g, ''))
            if (this.console_websocket != null) {
                this.console_websocket.send(item_input)
            } else {
                console.log("Warning, websocket is null!")
            }
        },
        consoleReceiveData: function(event){
            console.log(event)
        }
    },
    created(){
        var console_vue = this
        // Connect to WebSocket

        if (console_vue.console_websocket == null) {
            console.log("Connecting to Websocket!")
            console_vue.console_websocket = new WebSocket("ws://0.0.0.0:5000/api/ws")

            console_vue.console_websocket.onopen = function(){
                console_vue.console_websocket.send("Hello!")
            }
            console_vue.console_websocket.onmessage = function(evt){
                console_vue.consoleReceiveData(evt.data)
            }
        

        } else {
            console.log("Websocket Already Connected")
        }
    }
}
</script>