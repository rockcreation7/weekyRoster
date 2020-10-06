<template>
  <div class="list row">
    <md-list>
      <md-list-item
        class="list-group-item"
        v-for="(roster, index) in rosters"
        :key="index"
      >
        <div class="md-list-item-text">
          <span>{{ getDay(roster.date) }} {{ roster.date.slice(5, 10) }}</span>
          <span>{{ roster.upperStaff }} {{ roster.upperTime }}</span>
          <span>{{ roster.lowerStaff }} {{ roster.lowerTime }}</span>
        </div>
        <a>
          <md-button
            class="md-icon-button md-list-action"
            :href="'/roster/' + roster.date"
          >
            <md-icon class="md">edit</md-icon>
          </md-button>
        </a>
      </md-list-item>
    </md-list>
  </div>
</template>

<script>
import RosterDataService from "../services/RosterDataService";

export default {
  name: "roster-list",
  data() {
    return {
      rosters: [],
      shop: "nf",
    };
  },
  methods: {
    getDay(date) {
      switch (new Date(date).getDay()) {
        case 0:
          return "星期日";
        case 1:
          return "星期一";
        case 2:
          return "星期二";
        case 3:
          return "星期三";
        case 4:
          return "星期四";
        case 5:
          return "星期五";
        case 6:
          return "星期六";
      }
    },
    retrieveRosters() {
      RosterDataService.getAll(this.shop)
        .then((response) => {
          this.rosters = response.data;
          console.log(response.data);
        })
        .catch((e) => {
          this.$store.commit("errorMessage", e.message);
          console.log(e);
        });
    },
    changeShop(shop) {
      this.shop = shop;
      this.retrieveRosters();
    },
  },
  mounted() {
    this.retrieveRosters();
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