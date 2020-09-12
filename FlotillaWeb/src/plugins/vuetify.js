import Vue from 'vue';
import Vuetify from 'vuetify/lib';


import FilesIcon from "@/assets/vue-svg/hdd.vue";


Vue.use(Vuetify);

export default new Vuetify({
    icons:{
        values: {
            files: {
                component: FilesIcon
            }

        }
    }
});
