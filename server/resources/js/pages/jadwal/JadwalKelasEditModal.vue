<template>
<modal name="jadwal-kelas-edit" draggable=".card-header" :width="400" :height="400.6" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Edit : Hari {{hari}} & Mapel {{mapel}}</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('jadwal-kelas-edit')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-6">
        <label for="jam">Jam Mulai</label>
        <date-picker v-model="start" placeholder="Mulai" :minute-step="5" format="HH:mm" value-type="format" type="time"></date-picker>        
      </div>
      <div class="col-md-6">
        <label for="jam">Jam Akhir</label>
        <date-picker v-model="end" placeholder="Akhir" :minute-step="5" format="HH:mm" value-type="format" type="time"></date-picker>        
      </div>
      <div class="col-md-12" style="margin-bottom: 38px;">
        <label for="kelas">Guru</label>
        <v-select :options="guruData" label="nama" :reduce="nama => nama.id" v-model="guruList" placeholder="Pilih Guru" style="position: fixed;width: 27.8%;z-index: 999;"></v-select>
      </div>
      <div class="col-md-12">
        <label for="hari">Hari</label>
        <select class="form-control" @change="onHari($event)" v-model="hari">
          <option>Senin</option>
          <option>Selasa</option>
          <option>Rabu</option>
          <option>Kamis</option>
          <option>Jumat</option>
          <option>Sabtu</option>
          <option>Minggu</option>
        </select>  
      </div>  
      <div class="col-md-12">
        <label for="jam">Mapel</label> 
        <select v-model="jam" class="form-control">
          <option v-if="jamlist.length === this.$parent.maxpel" value="full">Jam Penuh</option>
          <option v-for="index in this.$parent.maxpel" v-if="!jamlist.includes(index)" :key="index" :value="index">Jam Ke {{index}}</option>
        </select>
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" @click="updatejadwal">Update Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'JadwalKelasEditModal',
  data () {
    return {
      id:'',              
      start:'',
      end:'',
      guruList:'',
      mapel:'',
      guruData: [],     
      hari:'',
      jam:'',      
      jamlist:[] 
    }
  },
  watch: {
    jamlist(newVal) {
      if(newVal === null){        
         this.jamlist = [0];
      }
    }
  },
  methods: {
    onHari(event){
      //remove variable select
      this.fetchJadwal()
      this.jam = 0;   
    },
    arrayRemove(arr, value) {
     return arr.filter(function(ele){
         return ele != value;
     });
    },
    fetchJadwal() {
    axios.get(window.location.origin + "/api/cekmapel?kelas="+this.$parent.kelasList+"&hari="+this.hari)
           .then(response => {
              this.jamlist = this.arrayRemove(response.data.data, this.jam);             
           }).catch(error => this.jamlist = [])
    },  
    fetchGuru() {
            axios.get(window.location.origin + "/api/alluser")
                .then(response => {
                    this.guruData = response.data.data
                }).catch(error => this.guruData = [])
    },  
    updatejadwal(){
      
      if(this.jam === "full"){
          this.$notify({
            group: 'notifikasi',                  
            title: 'Update Data Gagal',
            type: 'error',
            text: "Silahkan Pilih Jam Yang Tersedia"
          });
        return
      }

      axios.patch(window.location.origin+'/api/jadwal/'+this.id, {
        kelas_id: Number(this.$parent.kelasList), user_id: Number(this.guruList), hari: this.hari, jam:this.jam , jam_start: this.start ,jam_end:this.end
        })
        .then(response => {
          this.$parent.fetchJadwal();
          if(response.data.status == "success")
          { 
            //console.log('berhasil diupdate');
            this.$modal.hide('jadwal-kelas-edit')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Update Jadwal Sukses',
              type: 'success',
              text: 'Data Jadwal Berhasil Di Update'
            });
          }
        })
        .catch(error => {        
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        }) 

    },
    beforeOpen (event) {
      this.jam = event.params.wow.jam;
      this.hari = event.params.wow.hari;      
      this.fetchJadwal()
      this.fetchGuru()
      this.start = event.params.wow.jam_start;      
      this.end = event.params.wow.jam_end;
      this.guruList = event.params.wow.user_id;            
      this.id = event.params.wow.id;      
      this.mapel = event.params.wow.Users.mapel;
    },
    beforeClose (event) {
      //console.log("close")
      this.start = this.end = this.jamlist = this.guruList = this.hari = this.jam = '';
    }
   //end method
  }
}
</script>