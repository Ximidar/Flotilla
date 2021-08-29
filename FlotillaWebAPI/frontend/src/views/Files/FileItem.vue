<template>
    <v-expansion-panel @click="ClickEvent">
        <v-expansion-panel-header class="pa-0 ma-0">
            <v-hover v-slot:default="{ hover }">
                <v-toolbar class="toolbar-override" rounded='lg' :elevation="hover ? 12 : 5" :color="hover ? 'secondary' : 'primary'">
                    <v-icon class="basic-icon" v-if="File.FileType === 'file'">$vuetify.icons.solid_file</v-icon>
                    <v-icon  class="basic-icon" v-else>$vuetify.icons.solid_folder</v-icon>
                    <v-toolbar-title class="pl-5" >{{File.Name}}</v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-icon class="mx-1 basic-icon">$vuetify.icons.solid_info</v-icon>
                    <span>{{ ReadableSize }}</span>
                    <v-icon class="mx-1 basic-icon">$vuetify.icons.solid_carrot</v-icon>
                    <span>{{ FileDate }}</span>
                </v-toolbar>
            </v-hover>
        </v-expansion-panel-header>
        <v-expansion-panel-content class="pa-0 pt-3 ma-0">
            <v-row v-if="File.FileType === 'file'">
                <v-btn @click="playFile">
                    <v-icon class="pr-1">$vuetify.icons.regular_play_circle</v-icon>
                    <span>Play</span>
                    <v-snackbar v-model="snackbar"
                            centered
                            timeout="4000"
                            color="secondary"
                    >{{ snackbar_text }}</v-snackbar>
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn>
                    <v-icon class="pr-1">$vuetify.icons.solid_download</v-icon>
                    <span>Download</span>
                </v-btn>
                <v-btn>
                    <v-icon class="pr-1">$vuetify.icons.solid_skull</v-icon>
                    <span>Delete</span>
                </v-btn>
            </v-row>
        </v-expansion-panel-content>
    </v-expansion-panel>
</template>

<script>
import flotilla from "@/flotilla"
export default {
    name: 'FileItem',
    props:{
        File:{
            type: Object,
            default: {}
        }
    },
    mixins: [flotilla],
    data: function() {
        return {
            ReadableSize: this.HumanReadable(this.File.Size),
            FileDate: this.ConvertUnixTimestamp(this.File.UnixTime),
            snackbar: false,
            snackbar_text: "Playing " + this.File.Name
        }
    },
    methods:{
        HumanReadable: function (bytes) {
            var thresh = 1024;
            if(Math.abs(bytes) < thresh) {
                return bytes + ' B'
            }
            var units = ['kB','MB','GB','TB','PB','EB','ZB','YB']
            var u = -1
            do {
                bytes /= thresh
                ++u
            } while(Math.abs(bytes) >= thresh && u < units.length - 1)
            return bytes.toFixed(1)+' '+units[u]
        },
        ConvertUnixTimestamp: function(timestamp) {
            var a = new Date(timestamp * 1000)
            var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec']
            var year = a.getFullYear()
            var month = months[a.getMonth()]
            var date = a.getDate()
            var time = date + ' ' + month + ' ' + year 
            return time
        },
        ClickEvent: function() {
            this.$emit('clicked', this.File)
        },
        playFile: function() {
            this.snackbar = true
            this.flotSelectFile(this.File).then( response=>{
                console.log(response)
                // play file after it is selected
                this.flotPostAction("Play")
                console.log("Sent Play")
            })

        }
    },
    watch: {
        File: function(newval, oldval){
            this.ReadableSize = this.HumanReadable(newval.Size)
            this.FileDate = this.ConvertUnixTimestamp(this.File.UnixTime)

            console.log(File)
            console.log(this.ReadableSize)
        }
    }
}
</script>

<style scoped>

hover:hover{
    opacity: 0.75;
}

.basic-icon{
    width: 50px;
    height: 50px;
}

.iconsize{
  width: 25px;
  height: 25px;
  fill:white;
  padding-right: 10px;
  padding-left: 5px;
  display:inline-block;
}

.name{
    color: #859900;
    
}

.size{
    color:#2aa198;
}

.date{
    color: #6c71c4;
    
}
</style>