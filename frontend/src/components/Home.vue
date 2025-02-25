<template>
  <el-card class="box-card" :body-style="{ padding: '0px' }" data-wails-drag>
      <!-- Profile Section -->
      <div class="profile-section" style="--wails-draggable:drag">
        <img style="width: 32px" src="../assets/images/logo.png" alt="">
        <div class="user-info">
          <h3>Git Proxy</h3>
        </div>
      </div>

      <!-- Proxy List -->
      <div class="teams-list">
        <div class="team-item" v-for="team in teams" :key="team.id" :class="{ 'is-active': selectedTeam === team.id }" @click="selectTeam(team.id)">
          <el-icon><Position /></el-icon>
          <span class="team-name">{{ team.name }}</span>
          <el-icon v-if="selectedTeam === team.id" class="check-icon">
            <Select />
          </el-icon>
        </div>
      </div>

      <!-- Footer Actions -->
      <div class="footer-actions">
        <el-divider />
        <div class="action-item" @click="handleSettings">
          <el-icon><Setting /></el-icon>
          <span>Settings</span>
        </div>
        <el-divider />
        <div class="action-item" @click="handleAbout">
          <el-icon><Warning /></el-icon>
          <span>About</span>
        </div>
        <el-divider />
        <div class="action-item" @click="handleQuit">
          <el-icon><SwitchButton /></el-icon>
          <span>Quit</span>
        </div>
      </div>

      <!-- Setting drawer -->
      <el-drawer v-model="settingDrawer" :direction="direction" title="Settings" :show-close="false" size="80%" style="border-radius: 10px 10px 0 0">
        <el-form :model="form">
          <el-form-item label="HTTP&ensp;">
            <el-input v-model="form.httpAddr" placeholder="127.0.0.1:7890" ></el-input>
          </el-form-item>
          <el-form-item label="Socks5">
            <el-input v-model="form.socks5Addr" placeholder="127.0.0.1:80" ></el-input>
          </el-form-item>
        </el-form>
        <template #footer>
          <div style="flex: auto">
            <el-button round @click="resetSettings">Reset</el-button>
            <el-button round type="primary" @click="saveSettings">Save</el-button>
          </div>
        </template>
      </el-drawer>

      <!-- About drawer -->
      <el-drawer v-model="aboutDrawer" :direction="direction" :show-close="false" size="80%" style="border-radius: 10px 10px 0 0">
        <strong>Support</strong>
        <img style="width: 160px;border-radius: 12px" src="../assets/images/pay.png" alt="">
      </el-drawer>
    </el-card>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Select, Setting, Position, Warning, SwitchButton } from '@element-plus/icons-vue'
import {ElMessage} from 'element-plus'
import {SwitchMode, QueryConfig, SaveSettings, ResetSettings} from "../../wailsjs/go/main/App";

const direction = ref('btt')
const selectedTeam = ref('intercom')
const settingDrawer = ref(false)
const aboutDrawer = ref(false)

const form = reactive({
  httpAddr: "",
  socks5Addr: ""
});

const teams = [
  {
    id: 'default',
    name: 'No Proxy',
    initial: 'N',
    avatarClass: 'equals-avatar'
  },
  {
    id: 'http',
    name: 'Http Proxy',
    initial: 'H',
    avatarClass: 'stripe-avatar'
  },
  {
    id: 'socks',
    name: 'Socks Proxy',
    initial: 'S',
    avatarClass: 'intercom-avatar'
  }
]

const selectTeam = (teamId) => {
  selectedTeam.value = teamId
  // ElMessage.success(`Switched to ${teams.find(t => t.id === teamId).name}`)
  SwitchMode(teamId).then((result) => {
    ElMessage({
      message: `Use ${mode} Mode: `+result,
      type: "success",
    });
  })
}

const saveSettings = async () => {
  await SaveSettings(form.httpAddr, form.socks5Addr);
  ElMessage({
    message: "Saved Success",
    type: "success",
  });
  settingDrawer.value = false;
};

const resetSettings = async () => {
  await ResetSettings()
  ElMessage({
    message: "Reset Success",
    type: "success",
  });
  settingDrawer.value = false;
};

const handleSettings = async () => {
  settingDrawer.value = true
  // ElMessage.info('Opening settings')
  const result = await QueryConfig();
  form.httpAddr = result.httpAddr
  form.socks5Addr = result.socksAddr
};

const handleAbout = () => {
  aboutDrawer.value = true
}

const handleQuit = () => {
  window.runtime.Quit();
}
</script>

<style scoped>
.el-card{
   /* background: transparent; */
}

.profile-section {
  padding: 16px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid var(--el-border-color-lighter);
}

.user-info {
  margin-left: 12px;
}

.user-info h3 {
  margin: 0;
  font-size: 18px;
  color: var(--el-text-color-primary);
}

.user-info p {
  margin: 4px 0 0;
  font-size: 12px;
  color: var(--el-text-color-secondary);
}

.teams-list {
  padding: 8px 0;
}

.team-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  cursor: pointer;
  transition: background-color 0.3s;
  position: relative;
}

.team-item:hover {
  background-color: var(--el-fill-color-light);
}

.team-item.is-active {
  background-color: var(--el-fill-color);
}

.team-name {
  margin-left: 8px;  /* 图标与文字间距 */
  flex: 1;
  text-align: left;  /* 显式设置左对齐 */
  overflow: hidden;  /* 防止文本溢出 */
  white-space: nowrap; /* 单行显示 */
  text-overflow: ellipsis; /* 过长显示省略号 */
  font-size: 14px;
  color: var(--el-text-color-primary);
}

.check-icon {
  position: absolute;
  right: 16px;
  color: var(--el-color-success);
}

.create-avatar {
  border: 2px dashed var(--el-border-color);
  background-color: transparent;
  color: var(--el-text-color-secondary);
}

.equals-avatar {
  background-color: #ffaaa5;
  color: #fff;
}

.stripe-avatar {
  background-color: #635BFF;
  color: #fff;
}

.intercom-avatar {
  background-color: #60bf56;
  color: #fff;
}

.footer-actions {
  margin-top: 8px;
}

.action-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.action-item:hover {
  background-color: var(--el-fill-color-light);
}

.action-item .el-icon {
  margin-right: 8px;
  font-size: 16px;
  color: var(--el-text-color-secondary);
}

.action-item span {
  font-size: 14px;
  color: var(--el-text-color-primary);
}

:deep(.el-divider--horizontal) {
  margin: 0;
}
</style>