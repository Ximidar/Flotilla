<template>
    <v-row align="center" justify="center" class="px-14">
        <v-col>
            <v-col cols="12" >
                <v-select
                :items="connections"
                v-model="selectedPort"
                filled
                label="Console"
                ></v-select>
            </v-col>
            <v-col cols="12" >
                <v-select
                :items="baudRate"
                v-model="selectedBaud"
                filled
                label="Baud"
                ></v-select>
            </v-col>
            <v-col cols="12" >
                <v-row  align="center" justify="center">
                    <v-btn :disabled="connected" @click="commConnect" color="primary" class="mx-2">
                        <span>Connect</span>
                    </v-btn>
                    <v-btn :disabled="!connected" @click="commDisconnect" color="warning" class="mx-2">
                        <span>Disconnect</span>
                    </v-btn>
                </v-row>
                
            </v-col>
        </v-col>
    </v-row>
    
</template>

<script>
import flotilla from '@/flotilla'

export default {
    data() {
        return{
            connected: false,
            connections: ["/dev/tty0", "fakeprinter"],
            baudRate: [115200, 250000],
            selectedPort: "",
            selectedBaud: 115200,
            commStatus: {
                type: Object,
                default: {}
            }

        }
    },
    mixins: [flotilla],
    created(){
        // Get Comm Options
        this.flotGetCommOptions().then( commOptions =>{
            this.connections = []
            commOptions.ports.ports.forEach(port =>{
                this.connections.push(port.address)
            })

            this.baudRate = []
            commOptions.bauds.bauds.forEach(baud =>{
                this.baudRate.push(baud.speed)
            })
        })

        // Get Comm Status
        this.flotGetCommStatus().then( commStatus =>{
            this.commStatus = commStatus
            this.connected = commStatus.connected

            if (this.connected) {
                // select the port and baud
                this.selectedPort = commStatus.port
                this.selectedBaud = commStatus.baud

            } else {
                this.selectedPort = this.connections[0]
                this.selectedBaud = this.baudRate[0]
            }
        })

    },
    methods:{
        commConnect: function() {
            console.log("Connect Button pushed!")

            let ci = this.flotCreateCommInit(this.selectedPort, this.selectedBaud)
            console.log(ci)
            this.flotSendCommInit(ci).then( () =>{
                this.flotCommConnect().then( () =>{
                    console.log("Check if we are connected!")
                    this.connected = true
                })
            })
        },
        commDisconnect: function() {
            console.log("Disconnect Button pushed")
            this.flotCommDisconnect().then( () =>{
                console.log("Check if we are disconnected!")
                this.connected = false
            })
        }
    }

}
</script>