<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Absensi Siswa</h1>
        </div>
        <div class="row row-cards row-deck">
        	<div class="col-12">
        		<div class="card">
                    <div class="card-header">                       
                        <div class="col-2" style="padding-left: 0px;">
                        <date-picker v-model="tanggal" value-type="YYYY-MM-DD" format="DD/MM/YYYY" type="date" placeholder="Pilih Tanggal" @input="onTanggal()">                          
                        </date-picker>
                        </div>                       
                    <div class="card-options">
                        <select class="form-control" v-model="perPage" @change="onChange($event)" style="width: 60px;margin-right:10px;">
                            <option>5</option>
                            <option>10</option>
                            <option>15</option>
                            <option>20</option>
                        </select>
                        <div class="input-group" style="width:260px;">
                            <input type="text" class="form-control" v-model="search" v-on:keyup="searchEnter" placeholder="Cari Siswa">
                            <button class="btn bg-transparent" @click="refresh()" style="margin-left: -40px; z-index: 100;">
                                <i class="fe fe-x text-muted"></i>
                            </button>
                            <div class="input-group-append">
                                <button class="btn btn-primary" type="button" @click="cari()">
                                    <i class="fe fe-search"></i>
                                </button>
                            </div>
                        </div>
                    </div>
                    </div>
                    <div class="card-body" style="padding:0 !important;">
                      <div class="table-responsive">
                            <table class="table table-striped table-vcenter text-nowrap card-table">
                              <thead class="thead-dark">
                                <tr>
                                    <th class="text-center w-1">#</th>
                                    <th class="text-center w-1">
                                    <input type="checkbox" v-model="selectAll" style="top: 2px; position: relative;left: 8px;" number>
                                    </th>                                    
                                    <th class="text-center w-1" @click="sortByColumn('absens.tanggal')">
                                    Tanggal Absensi
                                    <span v-if="'absens.tanggal' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center w-1" @click="sortByColumn('absens.semester')">
                                    Semester
                                    <span v-if="'absens.semester' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center" @click="sortByColumn('absen_children.nama')">
                                    Nama
                                    <span v-if="'absen_children.nama' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center" @click="sortByColumn('absen_children.jenis')">
                                    Jenis
                                    <span v-if="'absen_children.jenis' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>                                      
                                    <th class="text-center w-1" @click="sortByColumn('absens.kelas')">
                                    Kelas
                                    <span v-if="'absens.kelas' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>                                    
                                    <th class="text-center">Absen</th>
                                    <th class="text-center"></th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr class="" v-if="tableData.length === 0">
                                  <td class="lead text-center" :colspan="9">No data found.</td>
                                </tr>
                                <tr v-for="(data, key1) in splitedList" :key="data.id">
                                  <td class="text-center">{{ serialNumber(key1) }}</td>
                                  <td class="text-center">
                                    <input type="checkbox" v-model="selected" :value="data.id" style="top: 2px; position: relative;left: 8px;" number>
                                  </td>
                                  <td class="text-center" v-if="data.tanggal">
                                   {{moment(data.tanggal).format('dddd')}} {{ data.tanggal | formatDate}}
                                  </td>
                                  <td class="text-center" v-if="data.semester">
                                   {{ data.semester }}                                  
                                  </td>
                                  <td class="text-center" v-if="data.nama">
                                   {{ data.nama }}
                                  </td>
                                  <td class="text-center" v-if="data.jenis">
                                   {{ data.jenis }}
                                  </td>                                                                    
                                  <td class="text-center" v-if="data.kelas">
                                   {{ data.kelas }}
                                  </td>                                                                     
                                  <td class="text-center">                                   
                                   <span class="badge" v-for="index in data.kode" style="color:#fff;margin-right:3px;border-radius:0px !important;" v-bind:style="{ backgroundColor: KodeColor(index)}">
                                   {{index.toUpperCase()}}                                   
                                   </span>                              
                                  </td>
                                  <td class="text-center">
                                    <button @click="changeAbsen(data.tanggal,data.siswa_id)" style="font-size:10px;font-weight: 600;line-height: 0;padding: 0.375rem 0.75rem" class="btn btn-primary" type="button">
                                      <i class="fe fe-edit-3" style="font-size: 12px;"></i>
                                      UBAH
                                    </button>  
                                  </td>
                                </tr>
                              </tbody>
                            </table>
                      </div>
                    </div>
                    <div class="card-footer">
                      <span v-if="pagination.total_record" style="position: absolute;bottom: 20px;"> &nbsp; <i>Total Data : {{ pagination.total_record }} entries.</i></span>
                        <nav v-if="pagination && tableData.length > 0">
                        <ul class="pagination float-right">
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(1)">
                                <i class="fe fe-chevrons-left"></i>
                              </a>
                            </li>
                            <li class="page-item" :class="{'disabled' : currentPage === 1}">
                              <a class="page-link" href="#" @click.prevent="changePage(currentPage - 1)"><i class="fe fe-chevron-left"></i></a>
                            </li>                            
                            <li v-for="page in pagesNumber" v-if="Math.abs(page - currentPage) < 3 || page == pagesNumber - 1 || page == 0" class="page-item" :class="{'active': page == pagination.page}">
                            <a href="javascript:void(0)" @click.prevent="changePage(page)" class="page-link">{{ page }}</a>
                            </li>                            
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(currentPage + 1)"><i class="fe fe-chevron-right"></i></a>
                            </li>
                            <li class="page-item" :class="{'disabled': currentPage === pagination.total_page}">
                            <a class="page-link" href="#" @click.prevent="changePage(pagination.total_page)"><i class="fe fe-chevrons-right"></i></a>
                            </li>
                        </ul>
                        </nav>
                    </div>
                </div>
        	</div>
        </div>        
        <absen-modal/>
    </div>
</template>
<script type="application/javascript">

import moment from "moment";
import "moment/locale/id";
import AbsenModal from './AbsenModal.vue'

export default {
  components : {  
    AbsenModal
  },
  data() {
    return {
      moment: moment,
      disable: true,    
      tanggal: null,    
      search:'',            
      tableData: [],
      url: window.location.origin,
      pagination:[],
      selected: [],
      offset: 4,
      currentPage: 1,
      perPage: '5',
      sortedColumn: 'absens.tanggal',
      order: 'desc'
    }
  },
  created: function() {
    moment.locale('id');    
    this.fetchData();
    this.RenderWebsocket();
  },
  watch: {
    selected(){
        if (this.selected.length===0){ this.disable = true} else { this.disable = false}
      }
  },
  computed: {
    splitedList() {
      let newArr = [...this.tableData]
      newArr.map(el => {
        return el.kode = el.kode.split(',')
      })
      return newArr
    },
    selectAll: {
            get: function () {
                return this.tableData ? this.selected.length == this.tableData.length : false;
            },
            set: function (value) {
                var selected = [];

                if (value) {
                    this.tableData.forEach(function (data) {
                        selected.push(data.id);
                    });
                }

                this.selected = selected;
            }
    },    
    /**
     * Get the pages number array for displaying in the pagination.
     * */
    pagesNumber() {
      if (this.pagination.offset == undefined || this.pagination.offset == null) {
        return []
      }
      let from = this.pagination.page - this.offset
      if (from < 1) {
        from = 1
      }
      let to = from + (this.offset * 2)
      if (to >= this.pagination.total_page) {
        to = this.pagination.total_page
      }
      let pagesArray = []
      for (let page = from; page <= to; page++) {
        pagesArray.push(page)
      }
      return pagesArray
    },
    /**
     * Get the total data displayed in the current page.
     * */
    totalData() {
      return (this.pagination.offset - this.pagination.total_record) + 1
    }
  },
  methods: {
    fetchData() {
      let dataFetchUrl = `${this.url}/api/absensi/harian?page=${this.currentPage}&tanggal=${this.tanggal}&search=${this.search}&order=${this.sortedColumn}&by=${this.order}&limit=${this.perPage}`
        axios.get(dataFetchUrl)
        .then(response => { 
          this.pagination = response.data
          this.tableData = response.data.records
        }).catch(error => this.tableData = [])
    },
    /**
     * Get the serial number.
     * @param key
     * */
    serialNumber(key) {
      return (this.currentPage - 1) * this.perPage + 1 + key
    },
    /**
     * Change the page.
     * @param pageNumber
     */
    changePage(pageNumber) {
      this.currentPage = pageNumber
      this.fetchData()
    },
    KodeColor(kode){
      if (kode === 'h'){return "green";}
      if (kode === 'i'){return "blue";}
      if (kode === 's'){return "#07ade3";}
      if (kode === 't'){return "orange";}
      if (kode === 'a'){return "red";}
    },
    onTanggal(){
      this.fetchData()
    },
    /** perpage ***/
    onChange(event) {
      this.currentPage = 1;
      this.fetchData();
    },    
    /** enter search **/
    searchEnter:function(e){
      if (e.keyCode === 13) {
        this.currentPage = 1;
        this.fetchData() 
      }
    },
    /** cari **/
    cari(){
      this.currentPage = 1;
      this.fetchData()
    },    
    /** refresh **/
    refresh(){
      this.search = "";
      this.fetchData();     
    },
    /** edit absen **/      
    changeAbsen(tanggal,siswa_id) {
      this.$modal.show('absen',{tgl:tanggal,siswaid:siswa_id});      
    },
    /** websocket **/
    RenderWebsocket: function() {
          this.socket = new WebSocket("ws://"+window.location.host+"/websocket/absen/ws");
          this.socket.onopen = () => {
            console.log("trigger connected");                     
            this.socket.onmessage = ({data}) => {          
             this.fetchData();
            };
          }
          this.socket2 = new WebSocket("ws://"+window.location.host+"/websocket/absenupdate/ws");
          this.socket2.onopen = () => {
            console.log("Realtime Update Konek");                     
              this.socket2.onmessage = ({data}) => {          
              this.fetchData();
            };
          }
    },
    /** multi hapus**/
    MultiDelete() {
      this.$modal.show('dialog', {
        title: 'Konfirmasi Multi Hapus',
        text: 'Apakah Anda benar-benar ingin menghapus?',
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
              axios.post(this.url+'/api/absensihapus',{selected:this.selected})
                .then(response => {                
                  if(response.data.status == "success")
                  { 
                    //console.log('berhasil dihapus');                       
                    this.fetchData();
                    this.$modal.hide('dialog');
                    this.$notify({
                      group: 'notifikasi',                  
                      title: 'Delete Sukses',
                      type: 'success',
                      text: 'Data Absensi Berhasil Di Hapus'
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
    /** pass variable to modal **/ 
    showDeleteDialog (data) {
      this.$modal.show('dialog', {
        title: 'Konfirmasi',
        text: 'Apakah Anda benar-benar ingin menghapus? Absensi : '+data.mapel_guru,
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
              axios.delete(this.url+'/api/absensi/'+data.id)
                .then(response => {                
                  if(response.data.status == "success")
                  { 
                    //console.log('berhasil dihapus');
                       
                    this.fetchData();
                    this.$modal.hide('dialog');
                    this.$notify({
                      group: 'notifikasi',                  
                      title: 'Delete Sukses',
                      type: 'success',
                      text: 'Data Absensi '+data.mapel_guru+' Berhasil Di Hapus'
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
    /**
     * Sort the data by column.
     * */
    sortByColumn(column) {
      if (column === this.sortedColumn) {
        this.order = (this.order === 'asc') ? 'desc' : 'asc'
      } else {
        this.sortedColumn = column
        this.order = 'asc'
      }
      this.fetchData()
    }
  },
  filters: {
    columnHead(value) {
      return value.split('_').join(' ').toUpperCase()
    }
  } 

//end script 
}
</script>