<template>
  <md-snackbar
    :md-position="position"
    :md-duration="isInfinity ? Infinity : duration"
    :md-active.sync="showSnackbar"
    md-persistent
  >
    <span>{{ error }}</span>
  </md-snackbar>
</template>
 
<script>
export default {
  data() {
    return {
      submitted: false,
      showSnackbar: false,
      position: "center",
      duration: 4000,
      isInfinity: false,
      error: "",
      /*  show: false,
      message: '',
      color: '' */
    };
  },

  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "errorMessage") {
        this.error = state.error;
        this.showSnackbar = true;
      }
    });
  },
};
</script>
