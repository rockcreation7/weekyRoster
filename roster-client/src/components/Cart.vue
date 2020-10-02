<template>
  <div class="list row">
    <md-list>
      <md-list-item
        class="list-group-item"
        v-for="(product, index) in cart"
        :key="index"
      >
        <div class="md-list-item-text">
          <span>{{ product.name }}</span>
          <span>{{ product.cartQty }}</span>
          <span>$ {{ product.price }}</span>
        </div>
        <!-- 
        <span @click="addToCart(product)">++</span>
        <span @click="rmFromCart(product)">--</span>
 -->
        <div>
          <div>
            <md-button
              @click="addToCart(product)"
              class="md-icon-button md-list-action"
            >
              <md-icon class="md">add</md-icon>
            </md-button>

            <md-button
              @click="rmFromCart(product)"
              class="md-icon-button md-list-action"
            >
              <md-icon class="md">remove</md-icon>
            </md-button>
          </div>
          <span>Code {{ product.code }}</span>
        </div>

        <md-button
          @click="dropFromCart(product.id)"
          class="md-icon-button md-list-action"
        >
          <md-icon class="md">delete</md-icon>
        </md-button>
      </md-list-item>
    </md-list>

    <div class="fucntionBar">
      <md-field>
        <md-input
          :type="getSearchType"
          v-model="search"
          placeholder="Input Price"
          inputmode="decimal"
        ></md-input>
      </md-field>
    </div>
    <div class="fucntionBar">
      <b>Cart Total : {{ cartTotal }}</b>
      <md-button class="md-raised clear" @click="clearCart()">Clear</md-button>
      <!--  <div>
        <button @click="setSerachType('byPrice')">by price</button>
        <button @click="setSerachType('byName')">by name</button>
      </div> -->
    </div>

    <md-list>
      <md-list-item
        class="list-group-item"
        v-for="(product, index) in filteredList"
        :key="index"
      >
        <div class="md-list-item-text">
          <span>{{ product.name }}</span>
          <span>$ {{ product.price }}</span>
          <p>Code {{ product.code }}</p>
        </div>
        <div class="md-list-item-text">
          <md-button class="md-raised md-primary" @click="addToCart(product)"
            >Add to Cart</md-button
          >
        </div>
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
      serachType: "",
    };
  },
  methods: {
    setSerachType(type) {
      this.searchType = type;
    },
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
    addToCart(product) {
      this.adjustCart(product, "add");
    },
    rmFromCart(product) {
      this.adjustCart(product, "rm");
    },
    dropFromCart(id) {
      this.cart = this.cart.filter((cartProduct) => cartProduct.id !== id);
    },
    clearCart() {
      this.cart = [];
    },
    adjustCart(product, operation) {
      console.log({ cart: this.cart });
      const cartProductIndex = this.cart.findIndex((cartProduct) => {
        return cartProduct.code === product.code;
      });

      if (this.cart.length === 0 || this.cart[cartProductIndex] === undefined) {
        product.cartQty = 1;
        this.cart.push(product);
        return;
      } else if (this.cart.length > 0) {
        if (operation === "add") {
          product.cartQty = this.cart[cartProductIndex].cartQty + 1;
        } else if (
          operation === "rm" &&
          this.cart[cartProductIndex].cartQty > 0
        ) {
          product.cartQty = this.cart[cartProductIndex].cartQty - 1;
        }
        this.$set(this.cart, cartProductIndex, product);
        return;
      }
    },
  },
  mounted() {
    this.retrieveProducts();
  },
  computed: {
    getSearchType() {
      switch (this.searchType) {
        case "byName":
          return "text";
        case "byPrice":
          return "number";
        default:
          return "number";
      }
    },
    filteredList() {
      return this.products.filter((product) => {
        if (this.search === "") {
          return false;
        }
        if (product.code) {
          // why "" still count as includes == true
          return (
            /* String(product.code).includes(this.search.toLowerCase()) ||
            product.name.toLowerCase().includes(this.search.toLowerCase()) || */
            String(product.price).includes(this.search.toLowerCase())
          );
        } else {
          return false;
        }
      });
    },
    cartTotal() {
      if (this.cart.length === 0) {
        return 0;
      }
      const calCartTotal = this.cart.reduce((accumulator, product) => {
        return accumulator + product.price * product.cartQty;
      }, 0);
      return Math.round(calCartTotal * 10) / 10;
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

.fucntionBar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}
</style>