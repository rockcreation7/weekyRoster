<template>
  <div class="list row">
    <input type="text" v-model="search" placeholder="Search Code/Name.." />
    
    <md-list>
      <md-list-item class="list-group-item" v-for="(product, index) in filteredList" :key="index">
        <span> {{product.name}}</span> 
        <span>$ {{product.price}}</span>
        <span> {{product.code}}</span>
        <a>
          <md-button class="md-icon-button md-list-action" :href="'/product/' + product.id">
            <md-icon class="md">edit</md-icon>
          </md-button>
        </a>
        <md-button @click="handleProductDelete(product.id)" class="md-icon-button md-list-action">
          <md-icon class="md">delete</md-icon>
        </md-button>
      </md-list-item>
    </md-list>
  </div>
</template>

<script>
import ProductDataService from "../services/ProductDataService";

export default {
  name: "Product-list",
  data() {
    return {
      search: "",
      products: [],
    };
  },
  methods: {
    retrieveProducts() {
      ProductDataService.getAll()
        .then((response) => {
          this.products = response.data;
          console.log(response.data);
        })
        .catch((e) => {
          this.$store.commit("errorMessage", e.message);
          console.log(e);
        });
    },
    handleProductDelete(id) {
      ProductDataService.delete(id)
        .then((response) => {
          console.log(response.data);
          console.log("successfully delete");
          this.retrieveProducts();
        })
        .catch((e) => {
          console.log(e);
          console.log("delete fail");
        });
    },
  },
  mounted() {
    this.retrieveProducts();
  },
  computed: {
    filteredList() {
      return this.products.filter((product) => {
        if (product.code) {
          return (
            String(product.code).includes(this.search.toLowerCase()) ||
            product.name.toLowerCase().includes(this.search.toLowerCase()) ||
            String(product.price).includes(this.search.toLowerCase())
          );
        } else {
          return false;
        }
      });
    },
  },
};
</script>

<style>
.list {
  max-width: 750px;
  margin: auto;
}

.md-list {
  width: 100%;
  display: inline-block;
  vertical-align: top;
  border: 1px solid rgba(#000, 0.12);
}
</style>