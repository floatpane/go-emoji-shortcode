import json

with open('/home/andrinoff/github/go-emoji-shortcode/emojis_raw.json', 'r', encoding='utf-8') as f:
    data = json.load(f)

seen = set()
out = []
for item in data:
    emoji = item.get('emoji')
    if not emoji:
        continue
    for alias in item.get('aliases', []):
        alias = alias.strip().lower()
        if alias and alias not in seen:
            seen.add(alias)
            out.append((alias, emoji))

with open('/home/andrinoff/github/go-emoji-shortcode/emojis.go', 'w', encoding='utf-8') as f:
    f.write('// Code generated from GitHub/gemoji emoji.json. DO NOT EDIT.\n\n')
    f.write('package shortcode\n\n')
    f.write('var byCode = map[string]string{\n')
    for alias, emoji in out:
        f.write('\t"%s": "%s",\n' % (alias.replace('\\', '\\\\').replace('"', '\\"'), emoji))
    f.write('}\n')

print('entries:', len(out))
