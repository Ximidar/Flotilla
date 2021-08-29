<template>

    <v-overlay :value="openOverlay"
    >
        <v-sheet
            color="primary"
            elevation="1"
        >
            <!-- file selector -->
            <v-container>
                <v-file-input
                    v-model="selectedFile"
                    chips
                    truncate-length="30"
                    accept="*.gcode"
                    label="GCODE Upload"
                    placeholder="Click Me"
                    prepend-icon=$vuetify.icons.solid_cube
                ></v-file-input>
            </v-container>
            <!-- buttons -->
            <v-container>
                <v-btn @click="fileUpload" color="primary" class="mx-2">
                    <span>Upload</span>
                </v-btn>
                <v-btn @click="closeOverlay" color="warning" class="mx-2">
                    <span>Cancel</span>
                </v-btn>
            </v-container>
        </v-sheet>

    </v-overlay>
    
</template>

<script>
import flotilla from "@/flotilla"
export default {
    

    name: "FileUploadOverlay",
    mixins: [flotilla],
    model: {
        prop: 'openOverlay',
        event: 'closeOverlay'
    },
    props: {
        openOverlay: {
            type: Boolean,
            default: false
        }
    },
    data(){
        return{
            selectedFile: null
        }
    },
    methods: {
        fileUpload: function(){
            this.flotUploadFile(this.selectedFile, "/etc/flotilla/gcode")
            this.$emit("closeOverlay", false)
        },
        closeOverlay: function(){
            console.log("Closing the overlay")
            this.$emit("closeOverlay", false)
        }
    }

    
}
</script>

<style scoped>

</style>