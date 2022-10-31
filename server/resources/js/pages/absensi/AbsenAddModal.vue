<template>
<modal name="absen-add" draggable=".card-header" :width="400" :height="533.4" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Non Jadwal</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('absen-add')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">
      <div class="col-md-6">
        <label for="tanggal">Tanggal</label>
        <date-picker v-model="tanggal" :disabled-date="notAfterToday" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal" @input="onTanggal()">                          
        </date-picker>
      </div>
      <div class="col-md-6" style="margin-bottom: 38px;">
        <label for="Kelas">Kelas</label>
        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas_id" placeholder="Pilih Kelas" @input="onKelas($event)" style="position: fixed;width: 13.8%;z-index: 999;">
        </v-select>         
      </div>
      <div class="col-md-12" style="margin-bottom: 38px;">
        <label for="mapel">Guru</label>
        <v-select :options="guruData" label="nama" :reduce="nama => nama.id" v-model="user_id" placeholder="Pilih Guru" @input="onGuru($event)" style="position: fixed;width: 27.8%;z-index: 990;">           
        </v-select>
      </div>
      <div class="col-md-6">
        <label for="mulai">Jam Masuk</label>
        <date-picker v-model="masuk" placeholder="Masuk Kelas" format="HH:mm" value-type="format" type="time" @input="onTanggal()"></date-picker>        
      </div>
      <div class="col-md-6">
        <label for="mulai">Jam Keluar</label>
        <date-picker v-model="keluar" placeholder="Keluar Kelas" format="HH:mm" value-type="format" type="time" @input="onTanggal()"></date-picker>        
      </div>
       <div class="col-md-12" style="margin-bottom: 38px;">
        <label for="note">Di Wakilkan <small style="font-size: 80%;font-weight: 700;">[DI KOSONGI JIKA TIDAK]</small></label>
        <v-select :disabled="disablepiket" :options="piketArray" label="nama" :reduce="nama => nama.id" v-model="piket_id" placeholder="Pilih Guru Piket" @input="onPiket()" style="position: fixed;width: 27.8%;z-index: 980;">
        </v-select>
      </div>      
      <div class="col-md-6">
        <label for="mulai">Jam Mulai Mapel</label>
        <date-picker v-model="start" placeholder="Mulai" :minute-step="5" format="HH:mm" value-type="format" type="time"></date-picker>        
      </div>
      <div class="col-md-6">
        <label for="akhir">Jam Akhir Mapel</label>
        <date-picker v-model="end" placeholder="Akhir" :minute-step="5" format="HH:mm" value-type="format" type="time"></date-picker>        
      </div>
      <div class="col-md-12">
        <label for="mapel">Mapel</label>
        <select v-model="jam" class="form-control" :disabled="disableJam">
          <option v-if="jamlist.length === maxpel" value="full">Jam Penuh</option>
          <option v-for="index in maxpel" v-if="!jamlist.includes(index)" :key="index" :value="index">Jam Ke {{index}}</option>
        </select>
      </div>     
    </div>
      <div class="float-right" style="margin-top:20px;">
        <button class="btn btn-primary" :disabled="disable" @click="adduser()">Input Data</button>
      </div> 
    </div>
  </div>   
  </modal>
</template>
<script type="application/javascript">
import moment from 'moment';
const today = new Date();
today.setHours(0, 0, 0, 0);

export default {
  name: 'AbsenAddModal',  
  data () {
    return {
      jamlist:[],
      guruData:[],
      kelasData:[],
      disablepiket:true,
      disableJam: true,
      disable: false,     
      tanggal:null,
      user_id:null,
      kelas_id:null,
      masuk:null,
      keluar:null,
      jam:'',
      kelas:'',
      mapel:'',
      mapel_guru:'',
      start:null,
      end:null,
      maxpel:'',
      piket_id:null,
      guru_piket:''
    }
  },
  computed: {
    piketArray: function() {
     return this.guruData.filter(item => item.id !== this.user_id )     
    }
  },
  watch: {
    jamlist(newVal) {
      if(newVal === null){        
         this.jamlist = [0];
      }
    },
    user_id(newVal){
      if(this.user_id !== null){        
        this.disablepiket = false;
      }else{
        this.disablepiket = true;
        this.guru_piket=null;
      }
    },
    piket_id(){
      if(this.piket_id == null){
        this.guru_piket = null;
      }
    },
    kelas_id(){
      if(this.kelas_id !== null && this.tanggal !== null){
         this.disableJam = false;
      }else{
        this.disableJam = true;
      }
    },
    tanggal() {
      if(this.tanggal !== null && this.kelas_id !== null){
         this.disableJam = false;
      }else{        
        this.disableJam = true;
      }
    }
  },
  methods: {
    getGuruAttr(){
      for (var key in this.guruData) {
          if (this.guruData.hasOwnProperty(key)) {
            if(this.user_id == this.guruData[key].id)
              {
                this.mapel = this.guruData[key].mapel;  
                this.mapel_guru = this.guruData[key].nama;  
              }
          }
      }
    },
    getPiketAttr(){
      for (var key in this.guruData) {
          if (this.guruData.hasOwnProperty(key)) {
            if(this.piket_id == this.guruData[key].id)
              {                
                this.guru_piket = this.guruData[key].nama;  
              }
          }
      }
    },
    getKelasAttr(){
      for (var key in this.kelasData) {
          if (this.kelasData.hasOwnProperty(key)) {
            if(this.kelas_id == this.kelasData[key].id)
              {                
                this.kelas = this.kelasData[key].nama;  
                this.maxpel = this.kelasData[key].maxpel;
              }
          }
      }
    },
    notAfterToday(date) {
      return date > today;
    },
    fetchJadwal() {
      axios.get(window.location.origin + "/api/absensi/cekmapel?kelas="+this.kelas_id+"&tanggal="+moment(this.tanggal).format('YYYY-MM-DD'))
           .then(response => {            
              this.jamlist = response.data.data              
           }).catch(error => this.jamlist = [])
    },
    fetchKelas(){
    axios.get(window.location.origin+"/api/allkelas")
      .then(response => {           
          this.kelasData = response.data.data                  
        }).catch(error => this.kelasData = [])
    },
    fetchGuru() {
      axios.get(window.location.origin + "/api/alluser")
           .then(response => {
                 this.guruData = response.data.data
      }).catch(error => this.guruData = [])
    },
    onKelas(event){
      this.getKelasAttr();
      this.fetchJadwal(); 
    },
    onGuru(event){
      this.getGuruAttr();
    },
    onPiket(event){
      this.getPiketAttr();
    },
    onTanggal(){
      this.fetchJadwal();      
    },
    adduser:function(event){
        if(this.kelas_id == null || this.start == null || this.end == null || this.masuk == null || this.keluar == null){
          this.$notify({
              group: 'notifikasi',                  
              title: 'Input Absensi Gagal',
              type: 'error',
              text: 'Periksa Kembali Input Absensi'
            });
          return
        }

        this.disable = true;      
        axios.post(window.location.origin+'/api/absensi/store', {
        tanggal:moment(this.tanggal+" "+this.masuk).format('YYYY-MM-DDTHH:mm'),keluar:moment(this.tanggal+" "+this.keluar).format('YYYY-MM-DDTHH:mm'),user_id:Number(this.user_id),kelas_id:Number(this.kelas_id),jam:Number(this.jam),kelas:this.kelas,mapel:this.mapel,mapelguru:this.mapel_guru,mapel_start:moment(this.tanggal+" "+this.start).format('YYYY-MM-DDTHH:mm'),mapel_end:moment(this.tanggal+" "+this.end).format('YYYY-MM-DDTHH:mm'),piketguru:this.guru_piket,piketid:Number(this.piket_id)
        })
        .then(response => {  
          this.$parent.fetchData();      
          if(response.data.status == "success")
          {             
            this.$modal.hide('absen-add')
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Absensi Sukses',
              type: 'success',
              text: 'Data Absensi '+this.mapel_guru+' Berhasil Di Tambahkan'
            });
          }
        })
        .catch(error => {
            this.disable = false;
            //console.log('gagal');
            this.$notify({
              group: 'notifikasi',                  
              title: 'Input Data Gagal',
              type: 'error',
              text: error.response.data.message
            });      
          
        })
    //-----    
    },
    beforeOpen (event) {
      this.fetchGuru();
      this.fetchKelas();
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")
      this.jamlist=[];
      this.guruData=[];
      this.kelasData=[];
      this.tanggal = this.keluar = this.kelas_id = this.start = this.end = this.user_id = null;
      this.jam = this.kelas = this.masuk = this.mapel = this.mapel_guru = this.maxpel = this.guru_piket = "";
    }
   //end method
  }
}
</script>