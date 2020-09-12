<template>
    <v-card class="pl-3">
        <span>{{ item_text }}</span>
    </v-card>
</template>

<script>
export default {
    props:{
        text: {
            type: String,
            default: ""
        }
    },
    data(){
        return {
            item_text: this.text
        }
    }
}
</script>