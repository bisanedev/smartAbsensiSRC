<template>
<modal name="jadwal-kelas-add" draggable=".card-header" :width="400" :height="400.6" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Jadwal Kelas</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('jadwal-kelas-add')">
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
        <v-select :options="guruData" label="nama" :reduce="nama => nama.id" v-model="guruList" placeholder="Pilih Guru" @input="onGuru($event)" style="position: fixed;width: 27.8%;z-index: 999;"></v-select>
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
        <select v-model="jam" class="form-control" :disabled="disable">
          <option v-if="jamlist.length === this.$parent.maxpel" value="full">Jam Penuh</option>
          <option v-for="index in this.$parent.maxpel" v-if="!jamlist.includes(index)" :key="index" :value="index">Jam Ke {{index}}</option>
        </select>
      </div>
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="tomboladd" @click="addjadwal">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
export default {
  name: 'JadwalKelasAddModal',
  data () {
    return {
      tomboladd:false,
      disable: true,           
      start:'',
      end:'',
      guruList:null,
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
      this.fetchJadwal()
      if(this.guruList !== null && this.hari !== ''){
        this.disable = false
      }else{        
        this.disable = true       
      }
    },

    onGuru(event){
      this.fetchJadwal()
      if (this.guruList !== null && this.hari !== '') {
          this.disable = false
      }else{
          this.disable = true
      }
    },

    fetchGuru() {
            axios.get(window.location.origin + "/api/alluser")
                .then(response => {
                    this.guruData = response.data.data
                }).catch(error => this.guruData = [])
    },

    fetchJadwal() {
      axios.get(window.location.origin + "/api/cekmapel?kelas="+this.$parent.kelasList+"&hari="+this.hari)
           .then(response => {
              this.jamlist = response.data.data              
           }).catch(error => this.jamlist = [])
    },    
    addjadwal(){

      if(this.jam === "full"){
          this.$notify({
            group: 'notifikasi',                  
            title: 'Input Data Gagal',
            type: 'error',
            text: "Silahkan Pilih Jam Yang Tersedia"
          });
        return
      }
      this.tomboladd = true;
      axios.post(window.location.origin+'/api/jadwal', {
        kelas_id: Number(this.$parent.kelasList), user_id: Number(this.guruList), hari: this.hari, jam:this.jam , jam_start: this.start ,jam_end:this.end
        })
        .then(response => {
          this.$parent.fetchJadwal();
          if(response.data.status == "success")
          {             
            //console.log('berhasil diupdate');
            this.$modal.hide('jadwal-kelas-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Jadwal Sukses',
              type: 'success',
              text: 'Data Jadwal Berhasil Di Tambahkan'
            });
          }
        })
        .catch(error => {               
            //console.log('gagal');
            this.tomboladd = false;
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        }) 

    },
    beforeOpen (event) {
      this.fetchGuru()
    },
    beforeClose (event) {
      this.tomboladd = false;
      //console.log("close")
      this.disable = true;
      this.guruList = null;
      this.start = this.end = this.jamlist  = this.hari = this.jam = '';   
    }
   //end method
  }
}
</script>