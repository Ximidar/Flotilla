<template>
<v-container>
    <v-card>
        <v-card-title>Flotilla Logo</v-card-title>
        <v-card-subtitle>Cool Logo Displayed Here</v-card-subtitle>
    </v-card>
    <v-row>
        <v-col>
            <PPCControl/>
            <v-card>
                <v-card-title>Connection</v-card-title>
                <v-card-subtitle>Status of the connection</v-card-subtitle>
                <v-btn color="primary" class="mx-2">Refresh</v-btn>
            </v-card>
        </v-col>
        <v-col>
            <v-card>
                <v-card-title>Temperature</v-card-title>
                <v-card-subtitle>Extruder and Bed Temperature</v-card-subtitle>
                <v-btn color="primary" class="mx-2">Refresh</v-btn>
            </v-card>
        </v-col>
    </v-row>
    <v-footer>
        <v-card-text>Flotilla version A0.1</v-card-text>
    </v-footer>
</v-container>
</template>

<script>
import flotilla from '@/flotilla'
import PPCControl from '@/views/Status/overview/PPCControl'
export default {
    name: "FlotillaOverview",
    components: {
        PPCControl
    },
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