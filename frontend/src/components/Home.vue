<template>
  <h2 id="result" class="result">Welcome to GitProxy</h2>
  <div class="mb-4">
    <el-button type="info" round @click="switchMode('default')">Default</el-button>
    <el-button type="primary" round @click="switchMode('http')">HTTP</el-button>
    <el-button type="success" round @click="switchMode('socks')">Socks</el-button>
    <el-button round @click="openSettings">Settings</el-button>
    <!-- Settings dialog -->
    <el-dialog
        v-model="settingsVisible"
        title="Settings"
        width="70%"
        center
    >
      <el-form :model="form" label-width="120px">
        <el-form-item label="HTTP">
          <el-input
              v-model="form.httpAddr"
              placeholder="127.0.0.1:7890"
          ></el-input>
        </el-form-item>
        <el-form-item label="Socks5">
          <el-input
              v-model="form.socks5Addr"
              placeholder="127.0.0.1:80"
          ></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetSettings">Reset</el-button>
          <el-button type="primary" @click="saveSettings">Save</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- Copyright information -->
    <el-footer class="footer">
      <p style="color: black">
        &copy; {{ currentYear }} YangYezhuang.
      </p>
    </el-footer>
  </div>
</template>

<script setup>
import {computed, reactive, ref} from "vue";
import {SwitchMode, QueryConfig, SaveSettings, ResetSettings} from "../../wailsjs/go/main/App";
import {ElMessage} from "element-plus";

const settingsVisible = ref(false);
const form = reactive({
  httpAddr: "",
  socks5Addr: "",
});
const currentYear = computed(() => new Date().getFullYear());

function switchMode(mode) {
  SwitchMode(mode).then((result) => {
    if (result == null || result === ''){
      ElMessage({
        message: `Please setting ${mode} addr`,
        type: "info",
      });
    }else {
      ElMessage({
        message: `Use ${mode} Mode: `+result,
        type: "success",
      });
    }
  });
}

const openSettings = async () => {
  settingsVisible.value = true;
  const result = await QueryConfig();
  form.httpAddr = result.httpAddr
  form.socks5Addr = result.socksAddr
};

const saveSettings = async () => {
  var result = await SaveSettings(form.httpAddr, form.socks5Addr);
  ElMessage({
    message: "Saved successfully",
    type: "success",
  });
  settingsVisible.value = false;
};

const resetSettings = async () => {
  ResetSettings()
  ElMessage({
    message: "Reset successfully",
    type: "success",
  });
  settingsVisible.value = false;
};
</script>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  color: #000;
}
.footer {
  background-color: #fff;
  width: 100%;
  text-align: center;
  padding: 16px 0;
  color: #4a5568;
  position: fixed;
  bottom: 0;
  left: 0;
}
</style>
