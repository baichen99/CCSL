<template>
  <list-view :get-list-method="GetAssignmentsList"></list-view>
</template>

<script>
import ListView from "@/components/ListView";
import {
  GetAssignmentsList,
  CreateAssignment,
  UpdateAssignment,
  DeleteAssignment
} from "@/api/assignments";
export default {
  name: "AssignmentList",
  components: {
    ListView
  },
  data() {
    return {
      GetAssignmentsList,
      CreateAssignment,
      UpdateAssignment,
      DeleteAssignment
    };
  }
};
</script>