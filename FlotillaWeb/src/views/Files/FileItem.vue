<template>
    <li>
        
        <FileIcon class="iconsize" v-if="FileType == 'file'"/>
        <FolderIcon class="iconsize" v-if="FileType == 'folder'"/>
        <div class="file-details">
            <div class="file-details name"><b>{{FileName}}</b></div>
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
        FileName:{
            default: 'Error',
            type: String
        },
        Path:{
            default: 'Error',
            type: String
        },
        FileType:{
            default: 'Error',
            type: String
        },
        Size:{
            default: -1,
            type: Number
        },
        UnixTime:{
            default: 0,
            type: Number
        },
        PreviousPath:{
            default: 'Error',
            type: String
        }
    },
    data: function() {
        return {
            ReadableSize: this.HumanReadable(this.Size),
            FileDate: this.ConvertUnixTimestamp(this.UnixTime)
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
        }
    },
    watch: {
        Size: function(newval, oldval){
            this.ReadableSize = HumanReadable(newval)
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