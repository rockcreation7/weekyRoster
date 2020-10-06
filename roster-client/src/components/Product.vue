<template>
  <div v-if="currentProduct" class="edit-form">
    <h4>Roster</h4>
    <form>
      <div class="form-group">
        <label for="Name">Name</label>
        <input type="text" class="form-control" id="Name" v-model="currentProduct.name" />
      </div>

      <div class="form-group">
        <label for="price">Price</label>
        <input type="text" class="form-control" id="Price" v-model="currentProduct.price" />
      </div>
      <div class="form-group">
        <label for="qty">Qty</label>
        <input type="text" class="form-control" id="Qty" v-model="currentProduct.qty" />
      </div>

      <div class="form-group">
        <label for="Cost">Cost</label>
        <input type="text" class="form-control" id="Cost" v-model="currentProduct.cost" />
      </div>
      <div class="form-group">
        <label for="code">Code</label>
        <input type="text" class="form-control" id="code" v-model="currentProduct.code" />
      </div>
      <div class="form-group">
        <label for="Catagory">Catagory</label>
        <input type="text" class="form-control" id="Catagory" v-model="currentProduct.catagory" />
      </div>

      <div class="form-group">
        <label for="Imgurl">Image Url</label>
        <input type="text" class="form-control" id="Imgurl" v-model="currentProduct.Imgurl" />
      </div>
    </form>
    <button type="submit" class="btn btn-primary" @click="updateProduct">Update</button>
    <span @click="handleProductDelete">
      <md-button class="md-icon-button md-list-action">
        <md-icon class="md">delete</md-icon>
      </md-button>
    </span>
  </div>

  <div v-else>
    <br />
    <p>Loading</p>
  </div>
</template>

<script>
import ProductDataService from "../services/ProductDataService";

export default {
  name: "Product",
  data() {
    return {
      currentProduct: null,
    };
  },
  methods: {
    getProduct(id) {
      ProductDataService.get(id)
        .then((response) => {
          console.log(response.data);
          this.currentProduct = response.data;
        })
        .catch((e) => {
          this.$store.commit("errorMessage", e.message);
          console.log(e);
        });
    },
    updateProduct() {
      ProductDataService.update(this.currentProduct.id, this.currentProduct)
        .then((response) => {
          console.log(response.data);
          this.$router.push("/products");
        })
        .catch((e) => {
          console.log(e);
        });
    },
    handleProductDelete() {
      ProductDataService.delete(this.currentProduct.id)
        .then(() => {
          console.log("successfully delete");
        })
        .catch(() => {
          console.log("delete fail");
        })
    },
  },
  mounted() {
    // this.message = "";
    this.getProduct(this.$route.params.id);
  },
};
</script>

<style>
.edit-form {
  max-width: 300px;
  margin: auto;
}
</style>