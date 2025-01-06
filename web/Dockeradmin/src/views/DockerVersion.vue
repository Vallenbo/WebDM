<template>
    <el-card>
        <div slot="header" class="clearfix">
            <span>版本信息</span>
        </div>
        <el-row>
            <el-col :span="12">
                <el-row>
                    <el-col :span="8">
                        <span>平台：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.Platform.Name }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>版本：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.Version }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>API版本：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.ApiVersion }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>MinAPI版本：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.MinAPIVersion }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>Git提交：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.GitCommit }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>Go版本：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.GoVersion }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>操作系统：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.Os }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>架构：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.Arch }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>内核版本：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.KernelVersion }}</span>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <span>构建时间：</span>
                    </el-col>
                    <el-col :span="16">
                        <span>{{ version.BuildTime }}</span>
                    </el-col>
                </el-row>
            </el-col>

        </el-row>
        <el-row>
            <el-col :span="16">
                <el-table :data="version.Components" style="width: 100%">
                    <el-table-column prop="Name" label="名称"></el-table-column>
                    <el-table-column prop="Version" label="版本"></el-table-column>
                    <el-table-column prop="Details.GitCommit" label="Git提交"></el-table-column>
                </el-table>
            </el-col>
        </el-row>
    </el-card>
</template>

<script>
export default {
    data() {
        return {
            version: {
                Platform: {
                    Name: 'Docker Engine - Community'
                },
                Components: [
                    {
                        Name: 'Engine',
                        Version: '23.0.2',
                        Details: {
                            ApiVersion: '1.42',
                            Arch: 'amd64',
                            BuildTime: '2023-03-27T16:16:31.000000000+00:00',
                            Experimental: 'false',
                            GitCommit: '219f21b',
                            GoVersion: 'go1.19.7',
                            KernelVersion: '3.10.0-1160.el7.x86_64',
                            MinAPIVersion: '1.12',
                            Os: 'linux'
                        }
                    },
                    {
                        Name: 'containerd',
                        Version: '1.6.19',
                        Details: {
                            GitCommit: '1e1ea6e986c6c86565bc33d52e34b81b3e2bc71f'
                        }
                    },
                    {
                        Name: 'runc',
                        Version: '1.1.4',
                        Details: {
                            GitCommit: 'v1.1.4-0-g5fd4c4d'
                        }
                    },
                    {
                        Name: 'docker-init',
                        Version: '0.19.0',
                        Details: {
                            GitCommit: 'de40ad0'
                        }
                    }
                ],
                Version: '23.0.2',
                ApiVersion: '1.42',
                MinAPIVersion: '1.12',
                GitCommit: '219f21b',
                GoVersion: 'go1.19.7',
                Os: 'linux',
                Arch: 'amd64',
                KernelVersion: '3.10.0-1160.el7.x86_64',
                BuildTime: '2023-03-27T16:16:31.000000000+00:00'
            }
        }
    }
}
</script>

