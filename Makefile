gohome:
	@if ! grep -q 'export PATH=$$PATH:$$HOME/gowork/bin' ~/.bashrc; then \
		echo 'export PATH=$$PATH:$$HOME/gowork/bin' >> ~/.bashrc; \
		echo "PATH 设置已添加到 ~/.bashrc"; \
		echo "请执行 'source ~/.bashrc' 使改动生效，或重新打开终端"; \
	else \
		echo "PATH 设置已存在于 ~/.bashrc 中，无需重复添加"; \
	fi