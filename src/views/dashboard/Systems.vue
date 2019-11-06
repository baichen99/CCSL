<template>
  <div class="app-container flex-column">
    <div class="table-toolbar">
      <el-button type="primary" plain @click="handleDump">
        备份数据库
        <i class="el-icon-plus el-icon--right" />
      </el-button>
    </div>
  </div>
</template>

<script>
import { GetDatabaseDump } from "@/api/systems";
export default {
  name: "Systems",
  methods: {
    handleDump() {
      GetDatabaseDump(true).then(res => {
        const url = window.URL.createObjectURL(new Blob([res]));
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", `${new Date().toISOString()}.tar.gz`);
        document.body.appendChild(link);
        link.click();
      });
    }
  }
};
</script>