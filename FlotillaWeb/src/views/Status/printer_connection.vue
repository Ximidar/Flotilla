<template>
    <v-row align="center" justify="center" class="px-14">
        <v-col>
            <v-col cols="12" >
                <v-select
                :items="connections"
                filled
                label="Console"
                ></v-select>
            </v-col>
            <v-col cols="12" >
                <v-select
                :items="baudRate"
                filled
                label="Baud"
                ></v-select>
            </v-col>
            <v-col cols="12" >
                <v-row  align="center" justify="center">
                    <v-btn color="primary" class="mx-2">
                        <span>Connect</span>
                    </v-btn>
                    <v-btn color="warning" class="mx-2">
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
            baudRate: ["115200", "250000"],

        }
    },
    mixins: [flotilla],
    created(){
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
    }
}
</script>