package data

import "github.com/mylxsw/eloquent/migrate"

func Migrate20240315Mix(m *migrate.Manager) {
	m.Schema("20240315-ddl").Create("models", func(builder *migrate.Builder) {
		builder.Increments("id")
		builder.Timestamps(0)
		builder.SoftDeletes("deleted_at", 0)

		builder.String("model_id", 100).Nullable(false).Comment("模型ID")
		builder.String("name", 100).Nullable(false).Comment("模型名称")
		builder.String("short_name", 100).Nullable(true).Comment("模型简称")
		builder.String("description", 255).Nullable(true).Comment("模型描述")
		builder.String("avatar_url", 255).Nullable(true).Comment("模型头像")
		builder.TinyInteger("status", false, true).Nullable(false).Comment("模型状态：0-禁用，1-启用")
		builder.String("version_min", 20).Nullable(true).Comment("最低版本")
		builder.String("version_max", 20).Nullable(true).Comment("最高版本")
		builder.Json("meta_json").Nullable(true).Comment("模型元信息，JSON 格式")
		builder.Json("providers_json").Nullable(false).Comment("模型提供商，JSON 格式")
	})

	m.Schema("20240315-ddl").Create("channels", func(builder *migrate.Builder) {
		builder.Increments("id")
		builder.Timestamps(0)
		builder.SoftDeletes("deleted_at", 0)

		builder.String("type", 100).Nullable(false).Comment("渠道类型")
		builder.String("name", 100).Nullable(false).Comment("渠道名称")
		builder.String("server", 255).Nullable(true).Comment("渠道 API 服务地址")
		builder.String("secret", 255).Comment("渠道密钥，多值的使用格式 APIKey|SecretKey")
		builder.Json("meta_json").Comment("其它配置信息，如是否使用代理等")
	})

	m.Schema("20240315-dml").Raw("models", func() []string {
		return []string{
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (1, '2024-03-15 17:03:11', null, 'gpt-3.5-turbo', 'GPT-3.5', null, '速度快，成本低', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt35.png-avatar', 1, null, null, '{"restricted": true, "max_context": 3500}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (2, '2024-03-15 17:03:11', null, 'gpt-3.5-turbo-16k', 'GPT-3.5 16K', null, '3.5 升级版，支持 16K 长文本', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt35.png-avatar', 1, null, null, '{"restricted": true, "max_context": 14000}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (3, '2024-03-15 17:03:11', null, 'gpt-4', 'GPT-4', null, '能力强，更精准', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt4.png-avatar', 1, null, null, '{"restricted": true, "max_context": 7500}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (4, '2024-03-15 17:03:11', null, 'gpt-4-turbo-preview', 'GPT-4 Turbo', null, '能力强，更精准', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt4-preview.png-avatar', 1, null, null, '{"restricted": true, "max_context": 123904}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (5, '2024-03-15 17:03:11', null, 'gpt-4-vision-preview', 'GPT-4V（视觉）', 'GPT-4V', '拥有视觉能力', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt4-preview.png-avatar', 1, '1.0.8', null, '{"vision": true, "restricted": true, "max_context": 123904}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (6, '2024-03-15 17:03:11', null, 'gpt-4-32k', 'GPT-4 32k', null, '基于 GPT-4，但是支持4倍的内容长度', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt4.png-avatar', 1, null, null, '{"restricted": true, "max_context": 28000}', '[{"name": "openai"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (7, '2024-03-15 17:03:11', null, 'claude-instant-1', 'Claude instant', 'Claude instant', 'Anthropic\'s fastest model, with strength in creative tasks. Features a context window of 9k tokens (around 7,000 words).', 'https://ssl.aicode.cc/ai-server/assets/avatar/anthropic-claude-instant.png-avatar', 1, '1.0.5', null, '{"restricted": true, "max_context": 123904}', '[{"name": "Anthropic"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (8, '2024-03-15 17:03:11', null, 'claude-2', 'Claude 2.1', 'Claude 2', 'Anthropic\'s most powerful model. Particularly good at creative writing.', 'https://ssl.aicode.cc/ai-server/assets/avatar/anthropic-claude-2.png-avatar', 1, '1.0.5', null, '{"restricted": true, "max_context": 195904}', '[{"name": "Anthropic"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (9, '2024-03-15 17:03:11', null, 'claude-3-haiku', 'Claude 3 Haiku', 'Claude Haiku', 'Anthropic\'s most powerful model, Fastest and most compact model for near-instant responsiveness.', 'https://ssl.aicode.cc/ai-server/assets/avatar/anthropic-claude-haiku-bg.png-avatar', 1, '1.0.5', null, '{"vision": true, "restricted": true, "max_context": 195904}', '[{"name": "Anthropic"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (10, '2024-03-15 17:03:11', null, 'claude-3-sonnet', 'Claude 3 Sonnet', 'Claude Sonnet', 'Anthropic\'s most powerful model, Ideal balance of intelligence and speed for enterprise workloads.', 'https://ssl.aicode.cc/ai-server/assets/avatar/anthropic-claude-sonnet-bg.png-avatar', 1, '1.0.5', null, '{"vision": true, "restricted": true, "max_context": 195904}', '[{"name": "Anthropic"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (11, '2024-03-15 17:03:11', null, 'claude-3-opus', 'Claude 3 Opus', 'Claude Opus', 'Anthropic\'s most powerful model for highly complex tasks.', 'https://ssl.aicode.cc/ai-server/assets/avatar/anthropic-claude-opus-bg.png-avatar', 1, '1.0.5', null, '{"vision": true, "restricted": true, "max_context": 195904}', '[{"name": "Anthropic"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (12, '2024-03-15 17:03:11', null, 'gemini-pro', 'Gemini Pro', null, 'Google 最新推出的 Gemini Pro 大语言模型', 'https://ssl.aicode.cc/ai-server/assets/avatar/gemini.png-avatar', 1, '1.0.5', null, '{"restricted": true, "max_context": 12000}', '[{"name": "google"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (13, '2024-03-15 17:03:11', null, 'gemini-pro-vision', 'Gemini Pro（视觉）', 'Gemini Pro V', 'Google 最新推出的 Gemini Pro 大语言模型，该版本为视觉版，支持图片输入，但是不支持多轮对话', 'https://ssl.aicode.cc/ai-server/assets/avatar/gemini.png-avatar', 1, '1.0.8', null, '{"vision": true, "restricted": true, "max_context": 30000}', '[{"name": "google"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (14, '2024-03-15 17:03:11', null, 'chat-3.5', 'Chat 3.5', null, '速度快，成本低', 'https://ssl.aicode.cc/ai-server/assets/avatar/nanxian.png-avatar', 1, '1.0.5', null, '{"max_context": 3500}', '[{"name": "openai", "model_rewrite": "gpt-3.5-turbo"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (15, '2024-03-15 17:03:11', null, 'chat-4', 'Chat 4.0', null, '能力强，更精准', 'https://ssl.aicode.cc/ai-server/assets/avatar/beichou.png-avatar', 1, '1.0.5', null, '{"max_context": 7500}', '[{"name": "openai", "model_rewrite": "gpt-4-turbo-preview"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (16, '2024-03-15 17:03:11', null, 'general', '星火大模型 v1.5', '星火 1.5', '科大讯飞研发的认知大模型，支持语言理解、知识问答、代码编写、逻辑推理、数学解题等多元能力，服务已内嵌联网搜索功能', 'https://ssl.aicode.cc/ai-server/assets/avatar/xfyun-v1.5.png-avatar', 1, '1.0.3', null, '{"max_context": 3500}', '[{"name": "讯飞星火"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (17, '2024-03-15 17:03:11', null, 'generalv2', '星火大模型 v2.0', '星火 2.0', '科大讯飞研发的认知大模型，V2.0 在 V1.5 基础上全面升级，并在代码、数学场景进行专项升级，服务已内嵌联网搜索、日期查询、天气查询、股票查询、诗词查询、字词理解等功能', 'https://ssl.aicode.cc/ai-server/assets/avatar/xfyun-v2.png-avatar', 1, '1.0.3', null, '{"max_context": 7500}', '[{"name": "讯飞星火"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (18, '2024-03-15 17:03:11', null, 'generalv3', '星火大模型 v3.0', '星火 3.0', '科大讯飞研发的认知大模型，V3.0 能力全面升级，在数学、代码、医疗、教育等场景进行了专项优化，让大模型更懂你所需', 'https://ssl.aicode.cc/ai-server/assets/avatar/xfyun-v3.png-avatar', 1, '1.0.3', null, '{"max_context": 7500}', '[{"name": "讯飞星火"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (19, '2024-03-15 17:03:11', null, 'generalv3.5', '星火大模型 v3.5', '星火 3.5', '科大讯飞研发的认知大模型，V3.0 能力全面升级，在数学、代码、医疗、教育等场景进行了专项优化，让大模型更懂你所需', 'https://ssl.aicode.cc/ai-server/assets/avatar/xfyun-v3.png-avatar', 1, '1.0.3', null, '{"max_context": 7500}', '[{"name": "讯飞星火"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (20, '2024-03-15 17:03:11', null, 'model_ernie_bot_turbo', '文心一言 Turbo', '文心 Turbo', '百度研发的知识增强大语言模型，中文名是文心一言，英文名是 ERNIE Bot，能够与人对话互动，回答问题，协助创作，高效便捷地帮助人们获取信息、知识和灵感', 'https://ssl.aicode.cc/ai-server/assets/avatar/wenxinyiyan-turbo.png-avatar', 1, '1.0.3', null, '{"max_context": 7000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (22, '2024-03-15 17:03:11', null, 'model_ernie_bot', '文心一言', null, '百度研发的知识增强大语言模型增强版，中文名是文心一言，英文名是 ERNIE Bot，能够与人对话互动，回答问题，协助创作，高效便捷地帮助人们获取信息、知识和灵感', 'https://ssl.aicode.cc/ai-server/assets/creative/wenxinyiyan.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (23, '2024-03-15 17:03:11', null, 'model_ernie_bot_4', '文心一言 4.0', '文心 4.0', 'ERNIE-Bot-4 是百度自行研发的大语言模型，覆盖海量中文数据，具有更强的对话问答、内容创作生成等能力', 'https://ssl.aicode.cc/ai-server/assets/avatar/wenxinyiyan-4.png-avatar', 1, '1.0.5', null, '{"max_context": 7500}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (24, '2024-03-15 17:03:11', null, 'model_badiu_llama2_70b', 'Llama 2 70B 英文版', 'Llama2 70B', '由 Meta AI 研发并开源，在编码、推理及知识应用等场景表现优秀，暂不支持中文输出', 'https://ssl.aicode.cc/ai-server/assets/avatar/llama2.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (25, '2024-03-15 17:03:11', null, 'model_baidu_llama2_13b', 'Llama 2 13B 英文版', 'Llama2 13B', '由 Meta AI 研发并开源，在编码、推理及知识应用等场景表现优秀，暂不支持中文输出', 'https://ssl.aicode.cc/ai-server/assets/avatar/llama2.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (26, '2024-03-15 17:03:11', null, 'model_baidu_llama2_7b_cn', 'Llama 2 7B 中文版', 'Llama2 7B', '由 Meta AI 研发并开源，在编码、推理及知识应用等场景表现优秀，当前版本是千帆团队的中文增强版本', 'https://ssl.aicode.cc/ai-server/assets/avatar/llama2-cn.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (27, '2024-03-15 17:03:11', null, 'model_baidu_llama2_13b_cn', 'Llama 2 13B 中文版', 'Llama2 13B', '由 Meta AI 研发并开源，在编码、推理及知识应用等场景表现优秀，当前版本是千帆团队的中文增强版本', 'https://ssl.aicode.cc/ai-server/assets/avatar/llama2-cn.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (28, '2024-03-15 17:03:11', null, 'model_baidu_chatglm2_6b_32k', 'ChatGLM2 6B', 'ChatGLM2', 'ChatGLM2-6B 是由智谱 AI 与清华 KEG 实验室发布的中英双语对话模型，具备强大的推理性能、效果、较低的部署门槛及更长的上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatglm.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (29, '2024-03-15 17:03:11', null, 'model_baidu_aquila_chat7b', 'AquilaChat 7B', 'AquilaChat', 'AquilaChat-7B 是由智源研究院研发，支持流畅的文本对话及多种语言类生成任务，通过定义可扩展的特殊指令规范', 'https://ssl.aicode.cc/ai-server/assets/avatar/aquila.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (30, '2024-03-15 17:03:11', null, 'model_baidu_bloomz_7b', 'BLOOMZ 7B', 'BLOOMZ', 'BLOOMZ-7B 是业内知名的⼤语⾔模型，由 BigScience 研发并开源，能够以46种语⾔和13种编程语⾔输出⽂本', 'https://ssl.aicode.cc/ai-server/assets/avatar/BLOOMZ.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (31, '2024-03-15 17:03:11', null, 'model_baidu_xuanyuan_70b', '轩辕 70B', null, '由度小满开发，基于Llama2-70B模型进行中文增强的金融行业大模型，通用能力显著提升，在CMMLU/CEVAL等各项榜单中排名前列；金融域任务超越领先通用模型，支持金融知识问答、金融计算、金融分析等各项任务', 'https://ssl.aicode.cc/ai-server/assets/avatar/xuanyuan.jpg-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (32, '2024-03-15 17:03:11', null, 'model_baidu_chat_law', 'ChatLaw', null, '由壹万卷公司与北大深研院研发的法律行业大模型，在开源版本基础上进行了进一步架构升级，融入了法律意图识别、法律关键词提取、CoT推理增强等模块，实现了效果提升，以满足法律问答、法条检索等应用需求', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatlaw.png-avatar', 1, '1.0.3', null, '{"max_context": 3000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (33, '2024-03-15 17:03:11', null, 'model_baidu_mixtral_8x7b_instruct', 'Mixtral-8x7B-Instruct', null, '由 Mistral AI 发布的首个高质量稀疏专家混合模型 (MOE)，模型由8个70亿参数专家模型组成，在多个基准测试中表现优于 Llama-2-70B 及 GPT3.5，能够处理 32K 上下文，在代码生成任务中表现尤为优异', 'https://ssl.aicode.cc/ai-server/assets/avatar/mixtral.jpg-avatar', 1, '1.0.3', null, '{"max_context": 30000}', '[{"name": "文心千帆"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (34, '2024-03-15 17:03:11', null, 'qwen-turbo', '通义千问 Turbo', '千问 Turbo', '通义千问超大规模语言模型，支持中文英文等不同语言输入', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-turbo.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (35, '2024-03-15 17:03:11', null, 'qwen-plus', '通义千问 Plus', '千问 Plus', '通义千问超大规模语言模型增强版，支持中文英文等不同语言输入', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-plus.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (36, '2024-03-15 17:03:11', null, 'qwen-max', '通义千问 Max', '千问 Max', '通义千问超大规模语言模型增强版，支持中文英文等不同语言输入', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-max.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (37, '2024-03-15 17:03:11', null, 'qwen-max-longcontext', '通义千问 Max+', '千问 Max+', '通义千问 Max Long Context 版本，支持 30K 上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-max-longcontext.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 25000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (38, '2024-03-15 17:03:11', null, 'qwen-vl-plus', '通义千问（视觉）', '千问 VL', '通义千问 VL 具备通用 OCR、视觉推理、中文文本理解基础能力，还能处理各种分辨率和规格的图像，甚至能“看图做题”', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-vlplus.jpeg-avatar', 1, '1.0.3', null, '{"vision": true, "max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (39, '2024-03-15 17:03:11', null, 'qwen-7b-chat', '通义千问 7B', '千问 7B', '通义千问 7B 是阿里云研发的通义千问大模型系列的 70 亿参数规模的模型，开源', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-osc-2.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (40, '2024-03-15 17:03:11', null, 'qwen-14b-chat', '通义千问 14B', '千问 14B', '通义千问 14B 是阿里云研发的通义千问大模型系列的 140 亿参数规模的模型，开源', 'https://ssl.aicode.cc/ai-server/assets/avatar/qwen-osc-1.jpeg-avatar', 1, '1.0.3', null, '{"max_context": 6000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (41, '2024-03-15 17:03:11', null, 'baichuan2-7b-chat-v1', '百川2 7B', null, '由百川智能研发的大语言模型，融合了意图理解、信息检索以及强化学习技术，结合有监督微调与人类意图对齐，在知识问答、文本创作领域表现突出', 'https://ssl.aicode.cc/ai-server/assets/avatar/baichuan.jpg-avatar', 1, '1.0.5', null, '{"max_context": 4000}', '[{"name": "灵积"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (42, '2024-03-15 17:03:11', null, 'Baichuan2-53B', '百川2 53B', null, '由百川智能研发的大语言模型，融合了意图理解、信息检索以及强化学习技术，结合有监督微调与人类意图对齐，在知识问答、文本创作领域表现突出', 'https://ssl.aicode.cc/ai-server/assets/avatar/baichuan.jpg-avatar', 1, '1.0.5', null, '{"max_context": 7500}', '[{"name": "百川"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (43, '2024-03-15 17:03:11', null, 'nova-ptc-xl-v1', '商汤日日新（大）', '日日新（大）', '商汤科技自主研发的超大规模语言模型，能够回答问题、创作文字，还能表达观点、撰写代码', 'https://ssl.aicode.cc/ai-server/assets/avatar/sensenova.png-avatar', 1, '1.0.3', null, '{"max_context": 2000}', '[{"name": "商汤日日新"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (44, '2024-03-15 17:03:11', null, 'nova-ptc-xs-v1', '商汤日日新（小）', '日日新（小）', '商汤科技自主研发的超大规模语言模型，能够回答问题、创作文字，还能表达观点、撰写代码', 'https://ssl.aicode.cc/ai-server/assets/avatar/sensenova.png-avatar', 1, '1.0.3', null, '{"max_context": 2000}', '[{"name": "商汤日日新"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (45, '2024-03-15 17:03:11', null, 'hyllm', '混元标准版', '混元', '由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力。这是标准版，当前已经标注为弃用，请直接使用混元标准版', 'https://ssl.aicode.cc/ai-server/assets/avatar/hunyuan.png-avatar', 1, '1.0.5', null, '{"max_context": 3000}', '[{"name": "腾讯"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (46, '2024-03-15 17:03:11', null, 'hyllm_pro', '混元高级版', '混元+', '由腾讯研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力', 'https://ssl.aicode.cc/ai-server/assets/avatar/hunyuan.png-avatar', 1, '1.0.5', null, '{"max_context": 3000}', '[{"name": "腾讯"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (47, '2024-03-15 17:03:11', null, '360GPT_S2_V9', '360智脑', '360智脑', '由 360 研发的大语言模型，拥有独特的语言理解能力，通过实时对话，解答疑惑、探索灵感，用AI技术帮人类打开智慧的大门', 'https://ssl.aicode.cc/ai-server/assets/avatar/gpt360.jpg-avatar', 1, '1.0.5', null, '{"max_context": 2000}', '[{"name": "360"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (48, '2024-03-15 17:03:11', null, 'chatglm_turbo', 'ChatGLM Turbo', null, '智谱 AI 发布的对话模型，具备强大的推理性能、效果、较低的部署门槛及更长的上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatglm.png-avatar', 1, '1.0.5', null, '{"max_context": 32000}', '[{"name": "oneapi"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (49, '2024-03-15 17:03:11', null, 'chatglm_pro', 'ChatGLM Pro', null, '智谱 AI 发布的对话模型，具备强大的推理性能、效果、较低的部署门槛及更长的上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatglm.png-avatar', 1, '1.0.5', null, '{"max_context": 8000}', '[{"name": "oneapi"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (50, '2024-03-15 17:03:11', null, 'chatglm_std', 'ChatGLM Std', null, '智谱 AI 发布的对话模型，具备强大的推理性能、效果、较低的部署门槛及更长的上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatglm.png-avatar', 1, '1.0.5', null, '{"max_context": 8000}', '[{"name": "oneapi"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (51, '2024-03-15 17:03:11', null, 'chatglm_lite', 'ChatGLM Lite', null, '智谱 AI 发布的对话模型，具备强大的推理性能、效果、较低的部署门槛及更长的上下文', 'https://ssl.aicode.cc/ai-server/assets/avatar/chatglm.png-avatar', 1, '1.0.5', null, '{"max_context": 8000}', '[{"name": "oneapi"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (52, '2024-03-15 17:03:11', null, '01-ai.yi-34b-chat', '零一万物 Yi 34B', 'Yi', '零一万物打造的开源大语言模型，在多项评测中全球领跑，MMLU 等评测取得了多项 SOTA 国际最佳性能指标表现，以更小模型尺寸评测超越 LLaMA2-70B、Falcon-180B 等大尺寸开源模型', 'https://ssl.aicode.cc/ai-server/assets/avatar/yi-01.png-avatar', 1, '1.0.5', null, '{"max_context": 4000}', '[{"name": "openrouter", "model_rewrite": "01-ai/yi-34b-chat"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (53, '2024-03-15 17:03:11', null, 'SkyChat-MegaVerse', '天工 MegaVerse', '天工', '昆仑万维研发的大语言模型，具备强大的中文创作能力，复杂语境下的逻辑推理能力，以及可靠的任务执行能力', 'https://ssl.aicode.cc/ai-server/assets/avatar/sky.png-avatar', 1, '1.0.5', null, '{"max_context": 4000}', '[{"name": "sky"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (54, '2024-03-15 17:03:11', null, 'glm-4', 'GLM-4', null, '智谱 AI 全自研的大语言模型，提供了更强大的问答和文本生成能力，适合于复杂的对话交互和深度内容创作设计的场景', 'https://ssl.aicode.cc/ai-server/assets/avatar/glm.png-avatar', 1, '1.0.5', null, '{"max_context": 8000}', '[{"name": "zhipu"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (55, '2024-03-15 17:03:11', null, 'glm-3-turbo', 'GLM-3 Turbo', 'GLM-3', '智谱 AI 全自研的大语言模型，适用于对知识量、推理能力、创造力要求较高的场景，比如广告文案、小说写作、知识类写作、代码生成等', 'https://ssl.aicode.cc/ai-server/assets/avatar/glm.png-avatar', 1, '1.0.5', null, '{"max_context": 8000}', '[{"name": "zhipu"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (56, '2024-03-15 17:03:11', null, 'moonshot-v1-8k', 'Moonshot 8K', null, 'AI 创业公司月之暗面推出的大语言模型', 'https://ssl.aicode.cc/ai-server/assets/avatar/moonshot.png-avatar', 1, '1.0.5', null, '{"max_context": 7000}', '[{"name": "moonshot"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (57, '2024-03-15 17:03:11', null, 'moonshot-v1-32k', 'Moonshot 32K', null, 'AI 创业公司月之暗面推出的大语言模型', 'https://ssl.aicode.cc/ai-server/assets/avatar/moonshot.png-avatar', 1, '1.0.5', null, '{"max_context": 30000}', '[{"name": "moonshot"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (58, '2024-03-15 17:03:11', null, 'moonshot-v1-128k', 'Moonshot 128K', null, 'AI 创业公司月之暗面推出的大语言模型', 'https://ssl.aicode.cc/ai-server/assets/avatar/moonshot.png-avatar', 1, '1.0.5', null, '{"max_context": 120000}', '[{"name": "moonshot"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (59, '2024-03-15 17:03:11', null, 'yi-34b-chat', 'Yi 34B', null, '零一万物打造的开源大语言模型，支持聊天、问答、对话、写作、翻译等功能', 'https://ssl.aicode.cc/ai-server/assets/avatar/yi-01.png-avatar', 1, '1.0.5', null, '{"max_context": 3000}', '[{"name": "oneapi", "model_rewrite": "yi-34b-chat-0205"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (60, '2024-03-15 17:03:11', null, 'yi-34b-chat-200k', 'Yi 34B 200K', null, '零一万物打造的开源大语言模型，增强了问答对话交互和深度内容创作能力', 'https://ssl.aicode.cc/ai-server/assets/avatar/yi-01.png-avatar', 1, '1.0.5', null, '{"max_context": 195904}', '[{"name": "oneapi"}]')`,
			`INSERT INTO models (id, created_at, updated_at, model_id, name, short_name, description, avatar_url, status, version_min, version_max, meta_json, providers_json) VALUES (61, '2024-03-15 17:03:11', null, 'yi-vl-plus', 'Yi VL', null, '零一万物打造的开源大语言模型，支持通用图片问答、图表理解、OCR、视觉推理', 'https://ssl.aicode.cc/ai-server/assets/avatar/yi-01.png-avatar', 1, '1.0.5', null, '{"vision": true, "max_context": 3000}', '[{"name": "oneapi"}]')`,
		}
	})
}
