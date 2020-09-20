<template>
  <div class="list row">
    <md-list>
      <md-list-item class="list-group-item" v-for="(product, index) in filteredList" :key="index">
        <span>{{product.name}}</span>
        <span>$ {{product.price}}</span>
        <span>Code {{product.code}}</span>
        <button @click="addToCart(product)">Add to Cart</button>
      </md-list-item>
    </md-list>
    <div>
      <input type="text" v-model="search" placeholder="Search Code/Name.." />
    </div>

    <md-list>
      <md-list-item class="list-group-item" v-for="(product, index) in cart" :key="index">
        <span>{{product.name}}</span>
        <!-- <span>Cost : {{product.cost}}</span> -->
        <!-- <span>Qty : {{product.qty}}</span>
          <span>Code : {{product.code}}</span>
        <span>Catagory : {{product.code}}</span>-->
        <span>{{product.cartQty}}</span>
        <span>$ {{product.price}}</span>
        <span>Code {{product.code}}</span>
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
      cart: [],
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
    addToCart(product) {
      console.log({ cart: this.cart });
      const cartProductIndex = this.cart.findIndex((cartProduct) => {
        return cartProduct.code === product.code;
      });

      if (this.cart.length === 0 || this.cart[cartProductIndex] === undefined) {
        product.cartQty = 1;
        this.cart.push(product);
        return;
      } else if (this.cart.length > 0) {
        product.cartQty = this.cart[cartProductIndex].cartQty + 1;
        this.$set(this.cart, cartProductIndex, product);
        return;
      }
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
            product.name.toLowerCase().includes(this.search.toLowerCase())
          );
        } else {
          console.log("here");
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