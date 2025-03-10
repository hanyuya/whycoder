### 设置清华源
```bash
# 删除 conda 默认源
conda config --remove-key channels 
# conda 换清华源
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/r
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/msys2
# conda 设置清华源第三方源
conda config --set custom_channels.auto https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/
```

### 设置中科大
```bash
# conda 中科大源
# 删除 conda 默认源
conda config --remove-key channels 
# conda 设置中科大源
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/main
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/r
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/msys2
# 第三方源
conda config --set custom_channels.conda-forge https://mirrors.ustc.edu.cn/anaconda/cloud
conda config --set custom_channels.pytorch https://mirrors.ustc.edu.cn/anaconda/cloud
```

### 查看路径是否配置
```bash
conda config --show channels
conda info
```

### 还原默认源
```bash
conda config --remove-key channels
```

### conda cmd
```bash
# 获取版本号
conda --version 或 conda -V

# 检查更新当前conda
conda update conda

# 查看当前存在哪些虚拟环境
conda env list 或 conda info -e

# Create env
# Create env with python 3.7 and pip
conda create --name whatwhale python=3.7 pip

# Activate env
# Activate by name
conda activate whatwhale

#Deactivate env
conda deactivate

#Removing an environment
#Remove environment by name
# remove env
conda env remove --name whatwhale

# verify
conda env list

#查看--安装--更新--删除包
conda list：
conda search package_name# 查询包
conda install package_name
conda install package_name=1.5.0
conda update package_name
conda remove package_name
```