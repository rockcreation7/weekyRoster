<template>
  <div v-if="currentRoster" class="edit-form">
    <h4>Roster</h4>
    <form>
      <div class="form-group">
        <label for="Date">Date</label>
        <input type="text" class="form-control" id="Date" v-model="currentRoster.date" />
      </div>

      <div class="form-group">
        <label for="upperstaff">Upper Staff</label>
        <input type="text" class="form-control" id="upperstaff" v-model="currentRoster.upperStaff" />
      </div>
      <div class="form-group">
        <label for="upperTime">Upper Time</label>
        <input type="text" class="form-control" id="upperTime" v-model="currentRoster.upperTime" />
      </div>

      <div class="form-group">
        <label for="lowerStaff">Lower Staff</label>
        <input type="text" class="form-control" id="lowerStaff" v-model="currentRoster.lowerStaff" />
      </div>
      <div class="form-group">
        <label for="lowerTime">Lower Time</label>
        <input type="text" class="form-control" id="lowerTime" v-model="currentRoster.lowerTime" />
      </div>

      <div class="form-group">
        <label for="customMessage">Custom Message</label>
        <input
          type="text"
          class="form-control"
          id="customMessage"
          v-model="currentRoster.customMessage"
        />
      </div>
    </form>
    <DataList/>
    <button type="submit" class="btn btn-primary" @click="updateRoster">Update</button>
  </div>

  <div v-else>
    <br />
    <p>Loading</p>
  </div>
</template>

<script>
import RosterDataService from "../services/RosterDataService";
import DataList from "../components/DataList";

export default {
  components: {
    DataList,
  },
  name: "Roster",
  data() {
    return {
      currentRoster: null,
    };
  },
  methods: {
    getRoster(date) {
      RosterDataService.get(date)
        .then((response) => {
          console.log(response.data);
          this.currentRoster = response.data;
        })
        .catch((e) => {
          console.log(e);
        });
    },
    updateRoster() {
      RosterDataService.update(this.currentRoster.date, this.currentRoster)
        .then((response) => {
          console.log(response.data);
          this.$router.push("/");
        })
        .catch((e) => {
          console.log(e);
        });
    },
    /* updatePublished(status) {
      var data = {
        id: this.currentTutorial.id,
        title: this.currentTutorial.title,
        description: this.currentTutorial.description,
        published: status
      };

      RosterDataService.update(this.currentTutorial.id, data)
        .then(response => {
          this.currentTutorial.published = status;
          console.log(response.data);
        })
        .catch(e => {
          console.log(e);
        });
    },


    deleteTutorial() {
      TutorialDataService.delete(this.currentTutorial.id)
        .then(response => {
          console.log(response.data);
          this.$router.push({ name: "tutorials" });
        })
        .catch(e => {
          console.log(e);
        });
    }
  */
  },
  mounted() {
    // this.message = "";
    this.getRoster(this.$route.params.date);
  },
};
</script>

<style>
.edit-form {
  max-width: 300px;
  margin: auto;
}
</style>