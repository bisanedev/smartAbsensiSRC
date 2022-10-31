<template>
<modal name="absen-jadwal" draggable=".card-header" :width="400" :height="533.4" @before-open="beforeOpen" @before-close="beforeClose">
  <div class="card">
    <div class="card-header card-header-modal">
     <h5 style="position: absolute; margin-top: 15px;">Input Jadwal Hari {{moment().format('dddd')}}</h5>
      <div class="card-options">
       <button class="close" @click="$modal.hide('absen-jadwal')">
          <i class="fe fe-x text-muted"></i>
       </button>
      </div>
    </div>
    <div class="card-body card-body-modal">
    <div class="form-row">     
      <div class="col-md-12" style="margin-bottom: 38px;">
        <label for="Kelas">Kelas</label>
        <v-select :options="kelasData" label="nama" :reduce="nama => nama.id" v-model="kelas_id" placeholder="Pilih Kelas" @input="onKelas($event)" style="position: fixed;width: 27.8%;z-index: 999;">
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
      <div class="col-md-12">        
        <div v-if="jadwalData.length == 0" class="jadwalabsen-null">Jadwal Data Kosong</div>
        <div v-else class="jadwalabsen">
          <div class="jadwalabsen-card" v-for="val in jadwalData" v-if="!jamlist.includes(val.jam)" :key="val.id" :class="[{change: val.id === pickJadwal}]" v-bind:style="{'background-color': val.Users.color}" @click="pilihjadwal(val)">
            
              <div v-if="val.Users.foto" class="avatar d-block foto" v-bind:style="{ 'background-image':'url(./storage/guru/'+ val.Users.foto +')' }">            
              </div>
              <div v-else-if="val.Users.jenis == 'Perempuan'" class="avatar d-block foto" style="background-image: url(./assets/images/cewek.png)">            
              </div>
              <div v-else class="avatar d-block foto" style="background-image: url(./assets/images/cowok.png)">            
              </div>
              <div class="jadwalabsen-konten">
              <div class="jadwalabsen-guru">{{val.Users.nama}} [{{val.Users.mapel}}]</div>
              <div class="jadwalabsen-waktu">[{{val.jam_start}} ~ {{val.jam_end}}]</div>
              </div>

          </div>
        </div>
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
import "moment/locale/id";
export default {
  name: 'AbsenJadwalModal',  
  data () {
    return {      
      moment:moment,
      jamlist:[],
      guruData:[],
      kelasData:[],
      jadwalData:[],
      pickJadwal:'',              
      disable: false,           
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
      piket_id:'',
      guru_piket:''
    }
  },
  computed: {
 
  },
  watch: {
    jamlist(newVal) {
      if(newVal === null){        
         this.jamlist = [0];
      }
    },
  },
  methods: {
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
    pilihjadwal(val){
      this.pickJadwal = val.id; 
      this.start = this.end = null;
      this.mapel = this.mapel_guru = this.user_id  = this.jam = "";
      this.mapel = val.Users.mapel
      this.mapel_guru = val.Users.nama
      this.user_id = val.Users.id
      this.start = val.jam_start
      this.end = val.jam_end
      this.jam = val.jam
      this.masuk = val.jam_start
      this.keluar = val.jam_end
    }, 
    cekJadwal() {
      axios.get(window.location.origin + "/api/absensi/cekmapel?kelas="+this.kelas_id+"&tanggal="+moment().format('YYYY-MM-DD'))
           .then(response => {            
              this.jamlist = response.data.data              
           }).catch(error => this.jamlist = [])
    },
    fetchJadwal(){
    axios.get(window.location.origin + "/api/jadwalabsen?kelas="+this.kelas_id+"&hari="+moment().format('dddd'))
           .then(response => {            
              this.jadwalData = response.data.data              
           }).catch(error => this.jadwalData = [])
    },
    fetchKelas(){
    axios.get(window.location.origin+"/api/allkelas")
      .then(response => {           
          this.kelasData = response.data.data                  
        }).catch(error => this.kelasData = [])
    },  
    onKelas(event){
      if(this.kelas_id != null){
        this.getKelasAttr();
        this.fetchJadwal();
        this.cekJadwal(); 
      }      
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
        var tgl = moment().format('YYYY-MM-DD');
        axios.post(window.location.origin+'/api/absensi/store', {
        tanggal:moment(tgl+" "+this.masuk).format('YYYY-MM-DDTHH:mm'),keluar:moment(tgl+" "+this.keluar).format('YYYY-MM-DDTHH:mm'),user_id:Number(this.user_id),kelas_id:Number(this.kelas_id),jam:Number(this.jam),kelas:this.kelas,mapel:this.mapel,mapelguru:this.mapel_guru,mapel_start:moment(tgl+" "+this.start).format('YYYY-MM-DDTHH:mm'),mapel_end:moment(tgl+" "+this.end).format('YYYY-MM-DDTHH:mm'),piketguru:this.guru_piket,piketid:Number(this.piket_id)
        })
        .then(response => {  
          this.$parent.fetchData();      
          if(response.data.status == "success")
          {             
            this.$modal.hide('absen-jadwal')
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
      this.fetchKelas();      
    },
    beforeClose (event) {
      this.disable = false;
      //console.log("close")
      this.jamlist=[];
      this.jadwalData=[];
      this.kelasData=[];
      this.keluar = this.kelas_id = this.start = this.end = this.user_id = null;
      this.pickJadwal = this.jam = this.kelas = this.masuk = this.mapel = this.mapel_guru = this.maxpel = this.guru_piket = "";
    }
   //end method
  }
}
</script>