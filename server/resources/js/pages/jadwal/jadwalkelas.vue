<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Jadwal Kelas</h1>
        </div>
        <div class="row row-cards row-deck">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <div class="col-2" style="padding-left: 0px;">
                        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelasList" placeholder="Pilih Kelas" @input="onKelas($event)"></v-select>
                        </div> 
                        <div class="card-options">
                            <button type="button" class="btn btn-primary" :disabled="disable" @click="showAddModal()">
                                <i class="fe fe-plus"></i>
                            </button>
                        </div>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                        <div v-if="kelasList === null" class="jadwal-null">
                            <h2 class="text-muted"> Silahkan Pilih Kelas </h2>
                        </div>
                        <div v-else class="timetable-week">
                            <div class="timetable-columns">
                                <div v-for="hari in week" class="timetable-column">
                                    <div class="timetable-column-header">{{hari}}</div>
                                    <div class="timetable-column-grid">
                                        <div v-for="index in maxpel" :key="index" class="pelajaran">
                                            <div v-for="data in jadwalData" :key="data.jam" v-if="data.hari === hari && index === data.jam && data.Users.id !== 0" v-bind:style="{ backgroundColor: data.Users.color}" class="color">
                                                <a type="button" @click="showDeleteDialog(data)" class="jadwalclose">
                                                    <i class="fe fe-x"></i>
                                                </a>                                                
                                                <div v-if="data.Users.foto" class="avatar d-block jdfoto" v-bind:style="{ 'background-image':'url(./storage/guru/'+ data.Users.foto +')' }">
                                                </div>
                                                <div v-else-if="data.Users.jenis == 'Perempuan'" class="avatar d-block jdfoto" style="background-image: url(./assets/images/cewek.png)">
                                                </div>
                                                <div v-else class="avatar d-block jdfoto" style="background-image: url(./assets/images/cowok.png)">
                                                </div> 
                                                <a type="button" @click="JadwalModal(data)" class="timetable-title-wrap">
                                                    <div class="timetable-name">{{data.Users.mapel}}</div>
                                                    <div class="timetable-time">
                                                    {{data.jam_start}} - {{data.jam_end}}
                                                    </div>
                                                </a>                                               
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div> 
                        </div>
                    </div>
                    <div class="card-footer" style="padding: 20px;"></div>
                </div>
            </div>
        </div>
        <jadwal-kelas-add-modal />
        <jadwal-kelas-edit-modal />
    </div>
</template>
<script type="application/javascript">

import JadwalKelasAddModal from './JadwalKelasAddModal.vue'
import JadwalKelasEditModal from './JadwalKelasEditModal.vue'

export default {
    components: {
        JadwalKelasEditModal,
        JadwalKelasAddModal
    },
    data() {
        return {
            disable: true,
            week: ['Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu', 'Minggu'],
            kelasList: null,
            maxpel: '',
            kelasData: [],
            jadwalData: []
        }
    },
    created: function() {
        this.fetchKelas()
    },
    methods: {
        fetchKelas() {
            axios.get(window.location.origin + "/api/allkelas")
                .then(response => {
                    this.kelasData = response.data.data
                }).catch(error => this.kelasData = [])
        },
        fetchJadwal() {
            axios.get(window.location.origin + "/api/jadwal?kelas=" + this.kelasList)
                .then(response => {
                    this.jadwalData = response.data.data
                }).catch(error => this.jadwalData = [])
        },
        /** Kelas ***/
        onKelas(event) {
            if (this.kelasList == null) {
                this.disable = true
            } else {
                this.disable = false
                this.fetchJadwal()
                for (var key in this.kelasData) {
                    if (this.kelasData.hasOwnProperty(key)) {
                        if (this.kelasList == this.kelasData[key].id) {
                            this.maxpel = this.kelasData[key].maxpel;
                        }
                    }
                }
            }
        },
        //delete
        showDeleteDialog(data) {
            this.$modal.show('dialog', {
                title: 'Konfirmasi',
                text: 'Hapus Jadwal ? Hari '+data.hari+', Mapel '+data.Users.mapel+", Jam "+data.jam_start+" - "+data.jam_end,
                buttons: [{
                        title: 'Cancel',
                        handler: () => {
                            this.$modal.hide('dialog')
                        }
                    },
                    {
                        title: 'Delete',
                        default: true,
                        handler: () => {
                            axios.delete(window.location.origin + '/api/jadwal/' + data.id)
                                .then(response => {
                                    if (response.data.status == "success") {
                                        //console.log('berhasil dihapus');
                                        this.fetchJadwal();
                                        this.$modal.hide('dialog');
                                        this.$notify({
                                            group: 'notifikasi',
                                            title: 'Delete Sukses',
                                            type: 'success',
                                            text: "Data Jadwal " + data.hari + " Mapel " + data.Users.mapel + " Jam "+data.jam_start+" - "+data.jam_end+" Berhasil Di Hapus"
                                        });
                                    }
                                })
                                .catch(error => {
                                    console.log(error);
                                })
                            //-------
                        }
                    }
                ]
            })
        },
        //insert data
        showAddModal(){
            this.$modal.show('jadwal-kelas-add');
        },
        //update    
        JadwalModal(data) {            
            this.$modal.show('jadwal-kelas-edit',{wow:data});
        }
    }


    //end script  
}
</script>