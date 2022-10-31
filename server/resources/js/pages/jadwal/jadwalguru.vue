<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Jadwal Guru</h1>
        </div>
        <div class="row row-cards row-deck">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <div class="col-4" style="padding-left: 0px;">
                            <v-select :options="guruData" label="nama" :reduce="nama => nama.id" v-model="guruList" placeholder="Pilih Guru" @input="onGuru()"></v-select>
                        </div>                                                                         
                        <div class="card-options">
                            <button type="button" class="btn btn-primary" :disabled="disable" @click="showAddModal()">
                                <i class="fe fe-plus"></i>
                            </button>                          
                        </div>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                        <div v-if="guruList === null" class="jadwal-null">
                            <h2 class="text-muted"> Silahkan Pilih Guru </h2>
                        </div>  
                        <div v-else class="timetable-list">
                            <div v-for="hari in week" class="timetable-day">
                                <div class="timetable-header">{{hari}}</div>                                

                                <div class="timetable-content">
                                    <div v-for="data in jadwalData" :key="data.jam" v-if="data.hari === hari" class="timetable-item">
                                        <span class="timetable-color" v-bind:style="{ backgroundColor: data.Users.color}"></span>
                                        <a type="button" @click="JadwalModal(data)" class="timetable-title">
                                            <span class="timetable-time">
                                            {{data.jam_start}} - {{data.jam_end}}
                                            </span>
                                            <span class="timetable-name">
                                            Mapel {{data.Users.mapel}} [<b>Jam Ke {{data.jam}}</b>] - Kelas [<b>{{data.Kelas.nama}}</b>]
                                            </span>
                                        </a>
                                        <a type="button" @click="showDeleteDialog(data)" class="float-right deljadguru">HAPUS <i class="fe fe-x-square"></i></a>
                                    </div>
                                </div>                                

                            </div>
                        </div>
                    </div>
                    <div class="card-footer" style="padding: 20px;"></div>
                </div>
            </div>
        </div>
        <jadwal-guru-add-modal />
        <jadwal-guru-edit-modal />        
    </div>
</template>
<script type="application/javascript">
import JadwalGuruAddModal from './JadwalGuruAddModal.vue'
import JadwalGuruEditModal from './JadwalGuruEditModal.vue'

export default {
    components: {
        JadwalGuruEditModal,
        JadwalGuruAddModal
    },
    data() {
        return {
            disable: true,
            week:['Senin','Selasa','Rabu','Kamis','Jumat','Sabtu','Minggu'],
            guruList: null,            
            guruData: [],
            jadwalData: []
        }
    },
    created: function() {
        this.fetchGuru()
    },
    methods: {    
        onGuru(){
            if (this.guruList === null) {
                this.disable = true                
            } else {                
                this.disable = false
                this.fetchJadwal()                
            } 
        },
        fetchGuru() {
            axios.get(window.location.origin + "/api/alluser")
                .then(response => {
                    this.guruData = response.data.data
                }).catch(error => this.guruData = [])
        },
        fetchJadwal(){
            axios.get(window.location.origin + "/api/jadwal?guru="+this.guruList)
            .then(response => {
                    this.jadwalData = response.data.data
            }).catch(error => this.jadwalData = [])
        },
        showAddModal(){
            this.$modal.show('jadwal-guru-add');
        },
        //delete
        showDeleteDialog (data) {
          this.$modal.show('dialog', {
            title: 'Konfirmasi',
            text: 'Hapus Jadwal ? Hari '+data.hari+', Kelas '+data.Kelas.nama+", Jam "+data.jam_start+" - "+data.jam_end,
            buttons: [
              {
                title: 'Cancel',            
                handler: () => {
                  this.$modal.hide('dialog')
                }
              },
              {
                title: 'Delete',
                default: true,
                handler: () => {                          
                  axios.delete(window.location.origin +'/api/jadwal/'+data.id)
                    .then(response => {                
                      if(response.data.status == "success")
                      { 
                        //console.log('berhasil dihapus');                           
                        this.fetchJadwal()
                        this.$modal.hide('dialog');
                        this.$notify({
                          group: 'notifikasi',                  
                          title: 'Delete Sukses',
                          type: 'success',
                          text: "Data Jadwal "+data.hari+" Kelas "+data.Kelas.nama+" Jam "+data.jam_start+" - "+data.jam_end+" Berhasil Di Hapus"
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
        //modal
        JadwalModal(data) {
            this.$modal.show('jadwal-guru-edit',{wow:data});            
        }
    }
//end script  
}
</script>