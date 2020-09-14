<template>
    <v-container >
        <v-row align="center" justify="start">
            <v-col cols="2">
                <v-row align="center" justify="center">
                    <v-col align="center" justify="stretch">
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_up</v-icon>
                        </v-btn>
                    </v-col>
                </v-row>
                <v-row align="center" justify="space-around">
                    <v-col align="center" justify="center">
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_left</v-icon>
                        </v-btn>
                    </v-col>
                    <v-col align="center" justify="center">
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_store_alt</v-icon>
                        </v-btn>
                    </v-col>
                    <v-col align="center" justify="center">
                    
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_right</v-icon>
                        </v-btn>
                    </v-col>
                </v-row>
                <v-row align="center" justify="center">
                    <v-col align="center" justify="stretch">
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_down</v-icon>
                        </v-btn>
                    </v-col>
                </v-row>
            </v-col>
            <v-col cols="1" align="center" justify="center">
                <v-row align="center" justify="center">
                    <v-container>
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_up</v-icon>
                        </v-btn>
                    </v-container>
                </v-row>
                <v-row align="stretch" justify="center">
                    <v-container>
                        <v-btn icon large>
                            <span style="font-size:24px">Z</span>
                        </v-btn>
                    </v-container>
                </v-row>
                <v-row align="center" justify="center">
                    <v-container>
                        <v-btn icon large>
                            <v-icon>$vuetify.icons.solid_arrow_alt_circle_down</v-icon>
                        </v-btn>
                    </v-container>
                </v-row>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="2" align="center" justify="center">
                <v-card outlined flat height="55px">
                    <v-chip class="ma-2" color="primary">X</v-chip>
                    <v-chip class="ma-2" color="primary">Y</v-chip>
                    <v-chip class="ma-2" color="primary">Z</v-chip>
                </v-card>
            </v-col>
            <v-col cols="1" align="start" justify="center">
                <v-text-field
                    label="Travel in mm"
                    outlined
                    placeholder="10"
                    height="55px"
                >
                </v-text-field>
            </v-col>
        </v-row>
    </v-container>
</template>


<style scoped>
.motor-column{
    flex-direction: column;
}
</style>