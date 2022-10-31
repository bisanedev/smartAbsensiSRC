<template>
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Absensi Mapel</h1>
        </div>
        <div class="row row-cards row-deck">
        	<div class="col-12">
        		<div class="card">
                    <div class="card-header">
                        <div class="item-action dropdown">
                          <button type="button" data-toggle="dropdown" class="btn btn-primary">
                            <i class="fe fe-plus"></i>
                          </button>
                          <div class="dropdown-menu">
                            <button type="button" @click="showAddModal()" class="dropdown-item" >
                              <i class="fe fe-help-circle"></i> Non Jadwal 
                            </button>
                            <button type="button" @click="showAddJadwalModal()" class="dropdown-item" >
                              <i class="fe fe-calendar"></i> Jadwal Hari ini
                            </button>
                          </div>
                        </div>                      
                        <div class="col-2" style="padding-left: 0px;margin-left: 25px;">
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
                            <input type="text" class="form-control" v-model="search" v-on:keyup="searchEnter" placeholder="Cari Pengajar">
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
                                    <th class="text-center w-1" @click="sortByColumn('tanggal')">
                                    Tanggal
                                    <span v-if="'tanggal' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center" @click="sortByColumn('semester')">
                                    Semester
                                    <span v-if="'semester' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center">
                                    Mulai ~ Selesai
                                    </th> 
                                    <th class="text-center" @click="sortByColumn('mapel_guru')">
                                    Pengajar
                                    <span v-if="'mapel_guru' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center w-1" @click="sortByColumn('kelas')">
                                    Kelas
                                    <span v-if="'kelas' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>
                                    <th class="text-center" @click="sortByColumn('piket_guru')">
                                    Guru Piket
                                    <span v-if="'piket_guru' === sortedColumn">
                                    <i v-if="order === 'asc' " class="fe fe-arrow-up"></i>
                                    <i v-else class="fe fe-arrow-down"></i>
                                    </span>
                                    </th>  
                                    <th class="text-center">
                                    Waktu
                                    </th> 
                                    <th class="text-center">
                                    </th>                                                                     
                                    <th class="text-center">
                                      <div class="item-action dropdown">
                                        <a href="javascript:void(0)" data-toggle="dropdown" class="icon" style="margin-left: 13px;color:#fff !important;" :class="{notactive: disable}">
                                            <i class="fe fe-grid"></i>
                                        </a>
                                        <div class="dropdown-menu dropdown-menu-right">
                                          <button type="button" @click="MultiDelete()" class="dropdown-item" >
                                            <i class="fe fe-trash"></i> Hapus 
                                          </button>
                                        </div>
                                      </div>
                                    </th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr class="" v-if="tableData.length === 0">
                                  <td class="lead text-center" :colspan="12">No data found.</td>
                                </tr>
                                <tr v-for="(data, key1) in tableData" :key="data.id">
                                  <td class="text-center" style="font-size: 13px;">{{ serialNumber(key1) }}</td>
                                  <td class="text-center">
                                    <input type="checkbox" v-model="selected" :value="data.id" style="top: 2px; position: relative;left: 8px;" number>
                                  </td>
                                  <td class="text-center" v-if="data.tanggal" style="font-size: 13px;">
                                   {{moment(data.tanggal).format('dddd')}} {{ data.tanggal | formatDate}}                                 
                                  </td>
                                  <td class="text-center" v-if="data.semester" style="font-size: 13px;">
                                   {{ data.semester }}                                  
                                  </td>
                                  <td class="text-center" style="font-size: 11px;" v-if="data.keluar !='0001-01-01T00:00:00Z'">            
                                    {{moment(data.tanggal).format('HH:mm')}} ~ {{moment(data.keluar).format('HH:mm')}}
                                    <b style="color: var(--mycolor);">
                                    [{{ moment.duration(moment(data.keluar).diff(moment(data.tanggal))).asMinutes() }} Menit]
                                    </b>                                   
                                  </td>
                                  <td v-else class="text-center" style="font-size: 11px;">
                                    {{moment(data.tanggal).format('HH:mm')}} ~ 00:00
                                    <b style="color: orange;">
                                    [0 Menit]
                                    </b>
                                  </td>
                                  <td class="text-center" v-if="data.mapel_guru" style="font-size: 13px;">
                                  {{ data.mapel_guru }} <b style="color: var(--mycolor);">[{{ data.mapel }}]</b>
                                  </td>
                                  <td class="text-center" v-if="data.kelas" style="font-size: 13px;">
                                   {{ data.kelas }}
                                  </td>
                                  <td class="text-center" v-if="data.piket_guru" style="font-size: 13px;">
                                   {{ data.piket_guru }}
                                  </td>
                                  <td class="text-center" v-else>
                                   -
                                  </td>
                                  <td class="text-center" style="font-size: 11px;">
                                    {{moment(data.mapel_start).format('HH:mm')}} ~ {{moment(data.mapel_end).format('HH:mm')}}
                                    <b style="color: var(--mycolor);">
                                    [{{ moment.duration(moment(data.mapel_end).diff(moment(data.mapel_start))).asMinutes() }} Menit] 
                                    </b>
                                  </td>
                                  <td class="text-center">
                                   <button @click="showSiswaModal(data)" class="btn btn-primary" type="button" style="font-size:10px;font-weight: 600;line-height: 0;padding: 0.375rem 0.75rem">
                                    <i class="fe fe-users" style="font-size: 12px;"></i> SISWA
                                   </button>                              
                                  </td>
                                  <td class="text-center">
                                      <div class="item-action dropdown">
                                        <a href="javascript:void(0)" data-toggle="dropdown" class="icon">
                                          <i class="fe fe-grid"></i>
                                        </a>
                                        <div class="dropdown-menu dropdown-menu-right">
                                        <button type="button" @click="showEditModal(data)" class="dropdown-item"><i class="fe fe-edit"></i> Edit </button> 
                                        <div class="dropdown-divider"></div>
                                        <button class="dropdown-item" @click="showDeleteDialog(data)"><i class="fe fe-trash"></i> Delete </a> </button>
                                        </div>
                                      </div>
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
        <absen-add-modal/>
        <absen-jadwal-modal/>
        <absen-edit-modal/>
        <siswa-modal/>
    </div>
</template>
<script type="application/javascript">

import AbsenAddModal from './AbsenAddModal.vue'
import AbsenEditModal from './AbsenEditModal.vue'
import SiswaModal from './SiswaModal.vue'
import AbsenJadwalModal from './AbsenJadwalModal.vue'

import moment from "moment";
import "moment/locale/id";

export default {
  components : {  
  AbsenAddModal,
  AbsenJadwalModal,
  AbsenEditModal,
  SiswaModal
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
      sortedColumn: 'tanggal',
      order: 'desc'
    }
  },
  created: function() {
    this.fetchData();
    this.RenderWebsocket();
  },
  watch: {
    selected(){
        if (this.selected.length===0){ this.disable = true} else { this.disable = false}
      }
  },
  computed: {
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
      let dataFetchUrl = `${this.url}/api/absensi?page=${this.currentPage}&tanggal=${this.tanggal}&search=${this.search}&order=${this.sortedColumn}&by=${this.order}&limit=${this.perPage}`
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
    /** edit modal **/
    showEditModal(data){      
      this.$modal.show('absen-edit',{wow:data});
    },
    showSiswaModal(data){
      this.$modal.show('siswa',{wow:data});
    },
    /** add modal **/    
    showAddModal(){
      this.$modal.show('absen-add');
    },
    showAddJadwalModal(){
      this.$modal.show('absen-jadwal');
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