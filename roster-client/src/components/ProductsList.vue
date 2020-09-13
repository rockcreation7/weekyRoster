<template>
  <div class="list row">
    <input type="text" v-model="search" placeholder="Search Code/Name.." />
    <md-list>
      <md-list-item class="list-group-item" v-for="(product, index) in filteredList" :key="index">
        <div class="md-list-item-text">
          <span>Name : {{product.name}}</span>
          <span>{{product.cost}}</span>
          <span>{{product.price}}</span>
          <span>{{product.qty}}</span>
          <span>Code : {{product.code}}</span>
        </div>
        <a>
          <md-button class="md-icon-button md-list-action" :href="'/product/' + product.id">
            <md-icon class="md">edit</md-icon>
          </md-button>
        </a>
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
          console.log(e);
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
          return (String(product.code).includes(this.search.toLowerCase())) || (product.name.toLowerCase().includes(this.search.toLowerCase()));
        } else {
          if (this.search === "") {
            return true;
          }
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