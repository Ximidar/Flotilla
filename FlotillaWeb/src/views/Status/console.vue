<template>
    <v-card dark class="pt-3">
        <flex>
            <v-col sm="6" cols="12">
                <v-subheader><b>Console Monitor</b></v-subheader>
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
                    rows="1"
                    row-height="24"
                    @keyup="onEnter"

                ></v-textarea>
                
            </v-col>
            
        </flex>
    </v-card>
</template>

<script>
import ConsoleItem from "@/views/Status/console_item.vue"
export default {
    name: 'console',
    components: {
        ConsoleItem
    },
    data(){
        return{
            connected: false,
            console_text: ["line1", "line2", "line3"],
            console_history: [],
            history_pos: -1,
            temp_input: "",
            console_input: "",
            auto_clear_single: true
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
        }
    }
}
</script>