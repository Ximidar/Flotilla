<template>
<v-container>
    <v-card>
        <v-card-title>Flotilla Logo</v-card-title>
        <v-card-subtitle>Cool Logo Displayed Here</v-card-subtitle>
    </v-card>
    <v-col>
        <v-row>
            <v-card>
                <v-card-title>File Playing</v-card-title>
                <v-card-subtitle>Status of the current playing file. Name / Progress</v-card-subtitle>
                <v-btn color="primary" class="mx-2">Pause</v-btn>
                <v-btn color="warning" class="mx-2">Cancel</v-btn>
            </v-card>
            <v-card>
                <v-card-title>Connection</v-card-title>
                <v-card-subtitle>Status of the connection</v-card-subtitle>
                <v-btn color="primary" class="mx-2">Refresh</v-btn>
            </v-card>
        </v-row>
        <v-row> 
            <v-card>
                <v-card-title>Temperature</v-card-title>
                <v-card-subtitle>Extruder and Bed Temperature</v-card-subtitle>
                <v-btn color="primary" class="mx-2">Refresh</v-btn>
            </v-card>
        </v-row>
    </v-col>   
    <v-footer>
        <v-card-text>Flotilla version A0.1</v-card-text>
    </v-footer>
</v-container>
</template>

<script>
import flotilla from '@/flotilla'
export default {
    name: "FlotillaOverview",
    components: {},
    mixins: [flotilla],
    data(){
        return {
            connectionStatus: {},

        }
    },
    created(){

    }
}
</script>

<style scoped>

</style>