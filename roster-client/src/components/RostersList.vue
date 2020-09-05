<template>
  <div class="list row">
    <md-list class="md-triple-line">
      <md-list-item class="list-group-item" v-for="(roster, index) in rosters" :key="index">
        <div class="md-list-item-text">
          <span>{{getDay(roster.date)}} {{roster.date.slice(5,10)}}</span>
          <span>{{roster.upperStaff}} {{roster.upperTime}}</span>
          <span>{{roster.lowerStaff}} {{roster.lowerTime}}</span>
        </div>
        <a>
          <md-button class="md-icon-button md-list-action" :href="'/roster/' + roster.date">
            <md-icon class="md">edit</md-icon>
          </md-button>
        </a>
      </md-list-item>

      <md-divider class="md-inset"></md-divider>
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
      currentRoster: null,
      currentIndex: -1,
      date: "",
    };
  },
  methods: {
    getDay(date) {
      console.log(date);
      console.log(new Date(date));
      console.log(new Date(date).getDay);
      switch (new Date(date).getDay()) {
        case 0:
          return "Sunday";
        case 1:
          return "Monday";
        case 2:
          return "Tuesday";
        case 3:
          return "Wednesday";
        case 4:
          return "Thursday";
        case 5:
          return "Friday";
        case 6:
          return "Saturday";
      }
    },
    retrieveRosters() {
      RosterDataService.getAll()
        .then((response) => {
          this.rosters = response.data;
          console.log(response.data);
        })
        .catch((e) => {
          console.log(e);
        });
    },

    refreshList() {
      this.retrieveRosters();
      this.currentRoster = null;
      this.currentIndex = -1;
    },

    setActiveRoster(Roster, index) {
      this.currentRoster = Roster;
      this.currentIndex = index;
    },

    /*     removeAllRosters() {
      RosterDataService.deleteAll()
        .then(response => {
          console.log(response.data);
          this.refreshList();
        })
        .catch(e => {
          console.log(e);
        });
    }, */

    searchDate() {
      RosterDataService.findByDate(this.Date)
        .then((response) => {
          this.rosters = response.data;
          console.log(response.data);
        })
        .catch((e) => {
          console.log(e);
        });
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