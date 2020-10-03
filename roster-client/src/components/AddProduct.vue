<template>
  <div class="submit-form">
    <div v-if="!submitted">
      <div class="form-group">
        <label for="name">Name</label>
        <input
          type="text"
          class="form-control"
          id="name"
          required
          v-model="product.name"
          name="name"
        />
      </div>

      <div class="form-group">
        <label for="qty">Qty</label>
        <input
          type="number"
          class="form-control"
          id="qty"
          required
          v-model="product.qty"
          name="qty"
        />
      </div>

      <div class="form-group">
        <label for="price">Price</label>
        <input
          type="number"
          class="form-control"
          id="Price"
          required
          v-model="product.price"
          name="price"
        />
      </div>
      <div class="form-group">
        <label for="cost">Cost</label>
        <input
          type="number"
          class="form-control"
          id="cost"
          required
          v-model="product.cost"
          name="cost"
        />
      </div>

      <div class="form-group">
        <label for="code">Code</label>
        <input
          type="number"
          class="form-control"
          id="code"
          required
          v-model="product.code"
          name="code"
        />
      </div>

      <div class="form-group">
        <label for="catagory">Catagory</label>
        <input
          class="form-control"
          id="catagory"
          required
          v-model="product.catagory"
          name="catagory"
        />
      </div>

      <div class="form-group">
        <label for="imgurl">Image Url</label>
        <input
          class="form-control"
          id="imgurl"
          required
          v-model="product.imgurl"
          name="imgurl"
        />
      </div>
      <DataList />
      <button @click="saveProduct" class="btn btn-success">Submit</button>
    </div>
    <div v-else>
      <h4>You submitted successfully!</h4>
      <button class="btn btn-success" @click="newProduct">Add Product</button>
    </div>

    <md-snackbar
      :md-position="position"
      :md-duration="isInfinity ? Infinity : duration"
      :md-active.sync="showSnackbar"
      md-persistent
    >
      <span>{{ error }}</span>
    </md-snackbar>
  </div>
</template>

<script>
import ProductDataService from "../services/ProductDataService";
export default {
  name: "add-product",
  data() {
    return {
      product: {
        name: "",
        price: "",
        cost: "",
        qty: "",
        code: "",
        catagory: "",
        imgurl: "",
      },
      submitted: false,
      showSnackbar: false,
      position: "center",
      duration: 4000,
      isInfinity: false,
      error: "",
    };
  },
  methods: {
    saveProduct() {
      var data = {
        name: this.product.name,
        code: this.product.code,
        price: this.product.price,
        cost: this.product.cost,
        qty: this.product.qty,
        catagory: this.product.catagory,
        imgurl: this.product.imgurl,
      };
      console.log(data);

      ProductDataService.create(data)
        .then((response) => {
          this.submitted = true;
          console.log(JSON.stringify(response.data));
        })
        .catch((e) => {
          this.showSnackbar = true;
          this.error = e;
          console.log(e);
        });
    },

    newProduct() {
      this.submitted = false;
      this.product = {};
    },
  },
};
</script>

<style>
.submit-form {
  max-width: 300px;
  margin: auto;
}
</style>