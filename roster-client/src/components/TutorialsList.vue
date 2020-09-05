<template>
  <div class="list row">
    <div class="col-md-8">
      <div class="input-group mb-3">
        <input type="text" class="form-control" placeholder="Search by date" v-model="title" />
        <div class="input-group-append">
          <button class="btn btn-outline-secondary" type="button" @click="searchRoster">Search</button>
        </div>
      </div>
    </div>
    <div class="col-md-6">
      <h4>Rosters List</h4>
      <ul class="list-group">
        <li
          class="list-group-item"
          :class="{ active: index == currentIndex }"
          v-for="(roster, index) in rosters"
          :key="index"
          @click="setActiveRoster(roster, index)"
        >{{ roster.date }}</li>
      </ul>
      <!-- 
      <button class="m-3 btn btn-sm btn-danger" @click="removeAllRosters">
        Remove All
      </button>-->
    </div>
    <div class="col-md-6">
      <div v-if="currentTutorial">
        <h4>Roster</h4>
        <div>
          <label>
            <strong>Date:</strong>
          </label>
          {{ currentTutorial.date }}
        </div>
        <div>
          <label>
            <strong>Upper Staff:</strong>
          </label>
          {{ currentTutorial.upperStaff }}
        </div>
        <a class="btn btn-warning" :href="'/roster/' + currentTutorial.id">Edit</a>
      </div>
      <div v-else>
        <br />
        <p>Please click on a Tutorial...</p>
      </div>
    </div>
  </div>
</template>

<script>
import TutorialDataService from "../services/TutorialDataService";

export default {
  name: "tutorials-list",
  data() {
    return {
      tutorials: [],
      currentTutorial: null,
      currentIndex: -1,
      title: "",
    };
  },
  methods: {
    retrieveTutorials() {
      TutorialDataService.getAll()
        .then((response) => {
          this.tutorials = response.data;
          console.log(response);
          console.log(response.data);
        })
        .catch((e) => {
          console.log(e);
        });
    },

    refreshList() {
      this.retrieveTutorials();
      this.currentTutorial = null;
      this.currentIndex = -1;
    },

    setActiveTutorial(tutorial, index) {
      this.currentTutorial = tutorial;
      this.currentIndex = index;
    },

    removeAllTutorials() {
      TutorialDataService.deleteAll()
        .then((response) => {
          console.log(response.data);
          this.refreshList();
        })
        .catch((e) => {
          console.log(e);
        });
    },

    searchTitle() {
      TutorialDataService.findByTitle(this.title)
        .then((response) => {
          this.tutorials = response.data;
          console.log(response.data);
        })
        .catch((e) => {
          console.log(e);
        });
    },
  },
  mounted() {
    this.retrieveTutorials();
  },
};
</script>

<style>
.list {
  text-align: left;
  max-width: 750px;
  margin: auto;
}
</style>