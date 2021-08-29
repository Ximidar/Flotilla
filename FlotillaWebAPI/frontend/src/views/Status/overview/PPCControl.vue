<template>
    <v-card>
        <v-card-subtitle>
            Status: {{currentStatus}}
        </v-card-subtitle>
        <v-container v-if="currentStatus != 'Idle'">
            <v-btn @click="toggleAction" color="primary" class="mx-2">{{ playpause[pppos] }}</v-btn>
            <v-btn @click="cancelAction" color="warning" class="mx-2">Cancel</v-btn>
        </v-container>
        <v-card-subtitle v-else class="pt-0 pb-0 ma-0">
            No Files Playing
        </v-card-subtitle>
    </v-card>
</template>

<script>
import flotilla from '@/flotilla'
export default {
    name: "PPCControl",
    components: {},
    mixins: [flotilla],
    data(){
        return {
            currentStatus: "Idle",
            playpause: ["Play", "Pause"],
            pppos: 1,
        }
    },
    methods: {
        toggleAction: function(){
            this.flotPostAction(this.ppText)
            this.pppos = (this.pppos == 1) ? 0 : 1
        },
        cancelAction: function(){
            this.flotPostAction("Cancel")
        }
    },
    created(){
        // check the status
        this.flotGetStatus().then( status =>{
            this.currentStatus = status
        })
    }
}
</script>

<style scoped>

</style>