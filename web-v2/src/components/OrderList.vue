<template>
  <div>
    <div>
      <b-form
        @button="searchOrders"
      >
        <b-form-group
          id="input-group-1"
          label="Enter the order information"
          label-for="input-1"
        >
          <b-form-input
            id="input-1"
            v-model="searchBox"
            type="searchBox"
            placeholder="order name or product name"
            required
          ></b-form-input>
        </b-form-group>

        <b-button
          v-on:click="searchOrders(searchBox)"
          variant="primary"
        >Search</b-button>
        <!-- <b-button variant="primary">Search</b-button> -->
      </b-form>
    </div>

    <b-form @button="filterOrders">
      <label for="datepicker-start">Choose a start date</label>
      <b-form-datepicker
        id="datepicker-start"
        v-model="startDate"
        type="date"
        class="mb-2"
      ></b-form-datepicker>

      <label for="datepicker-end">Choose an end date</label>
      <b-form-datepicker
        id="datepicker-end"
        v-model="endDate"
        type="date"
        class="mb-2"
      ></b-form-datepicker>
      <b-button
        v-on:click="filterOrders"
        variant="primary"
      >Filter</b-button>
    </b-form>

    <p>Total amount: <b> ${{ totalAmount }} </b> </p>

    <div
      v-if="!hasOrders"
      style="text-align: center"
    ><br><br> NO DATA </div>
    <div v-if="hasOrders">
      <b-table
        :items="orders"
        :fields="ordersInfo"
        striped
        hover
        :current-page=1
        :per-page="perPage"
      >
        <template v-slot:cell(deliveredAmount)="data">
          <div v-if="data.value === '' ">
            <b> $0 </b>
          </div>
          <div v-else>
            <b> ${{data.value}}</b>
          </div>
        </template>

        <template v-slot:cell(totalAmount)="data">
          <div v-if="data.value === '' ">
            <b> $0 </b>
          </div>
          <div v-else>
            <b> ${{data.value}}</b>
          </div>
        </template>
      </b-table>

      <p class="mt-3"> Total Orders: {{ totalOrders }} </p>

      <b-row>
        <b-col
          md="6"
          class="my-1"
        >
          <b-pagination
            :total-rows="totalOrders"
            :per-page="perPage"
            v-model="currentPage"
            class="my-0"
          />
        </b-col>
        <b-col
          md="6"
          class="my-1"
        >
        </b-col>
      </b-row>
    </div>
  </div>
</template>


<script>
  export default {
    data() {
      return {
        items: [
          { id: 1, first_name: 'Fred', last_name: 'Flintstone' },
        ],
      }
    },
  }
</script>


<script>
import ordersJson from "../data/orders.json";

export default {
  data() {
    return {
      serverApiUrl: BaseConfig.ServerApiUrl,

      searchBox: "",
      startDate: '',
      endDate: '',
      totalAmount: 0,
      totalOrders: 0,
      orders: [],
      ordersInfo: [
        {
          key: "orderName",
          sortable: true,
          label: "Order Name"
        },
        {
          key: "companyName",
          sortable: true,
          label: "Customer Company",
          class: "text-right options-column"
        },
        {
          key: "customerName",
          sortable: true,
          label: "Customer Name",
          class: "text-right options-column"
        },
        {
          key: "orderData",
          sortable: true,
          label: "Order Data",
          class: "text-right options-column"
        },
        {
          key: "deliveredAmount",
          sortable: true,
          label: "Delivered Amount",
          class: "text-right options-column"
        },
        {
          key: "totalAmount",
          sortable: true,
          label: "Total Amount",
          class: "text-right options-column"
        },
      ],
        queryParams: {
            sort: [],
            filters: {
                startDate: '',
                endDate: '',
            },
            global_search: "",
            per_page: 10,
            page: 1,
        },

      selectAll: false,
      records: [],
      orders:[],
      perPage: 5,
      currentPage: 1,
    };
  },
  computed: {
    hasOrders() {
      return this.totalOrders > 0;
    },
    totalOrders() {
      return this.totalOrders;
    },
  },
  watch: {
    currentPage: function (val) {
      this.queryParams.page = val
      this.fetchData();
    },
  },
  methods: {
    clearnParams: function () {
        this.queryParams.filters.startDate = ''
        this.queryParams.filters.endDate = ''
        this.queryParams.global_search = ''
    },
    searchOrders: function (event) {
        this.clearnParams()
        this.queryParams.global_search = event
        this.fetchData();
      },
    filterOrders() {
        this.clearnParams()
        this.queryParams.filters.startDate = this.startDate
        this.queryParams.filters.endDate = this.endDate
        this.fetchData();
      },
    fetchData() {
        let self = this;
        axios.get(this.serverApiUrl + '/api/v1/orders', {
            params: {
                "StartOrderDate": this.queryParams.filters.startDate,
                "EndOrderDate": this.queryParams.filters.endDate,
                "SearchOrder": this.queryParams.global_search,
                "PageNum": this.queryParams.page
            }
        })
        .then(function(response) {
            // console.log(response);
            self.orders = response.data.data.lists;
            self.totalOrders = response.data.data.total;
            self.totalAmount =  response.data.data.totalAmount
        })
        .catch(function(error) {
            console.log(error);
        });
    }
  },
  components: {},
  mounted() {
    this.fetchData();
    // var vm = this;
    // setTimeout(function() {
    //   vm.orders = ordersJson.data.lists;
    //   vm.totalOrders = ordersJson.data.total;
    // }, 1000);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
