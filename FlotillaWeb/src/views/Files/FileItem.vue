<template>
    <li v-on:click="ClickEvent">
        
        <FileIcon class="iconsize" v-if="File.FileType === 'file'"/>
        <FolderIcon class="iconsize" v-else/>
        <div class="file-details">
            <div class="file-details name"><b>{{File.Name}}</b></div>
            <div class="file-details size">{{ ReadableSize }}</div> 
            <div class="file-details date">{{ FileDate }}</div>
        </div>
          
    </li>
</template>

<script>
import FileIcon from "@/assets/svg/solid/file.svg"
import FolderIcon from "@/assets/svg/solid/folder.svg"
export default {
    name: 'FileItem',
    components:{
        FileIcon,
        FolderIcon
    },
    props:{
        File:{
            type: Object,
            default: {}
        }
    },
    data: function() {
        return {
            ReadableSize: this.HumanReadable(this.File.Size),
            FileDate: this.ConvertUnixTimestamp(this.File.UnixTime)
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

li:hover{
    opacity: 0.75;
}

li{
  min-height: 50px;
  max-height: 50px;
  width: 100%;
  margin: 0 0 5px 0;
  text-align: left;
  cursor: default;
}

.iconsize{
  width: 25px;
  height: 25px;
  fill:white;
  padding-right: 10px;
  padding-left: 5px;
  display:inline-block;
}

.file-details{
    display:inline-block;
    transform: translate(0%, -20%);
    padding-right: 10px;
    
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