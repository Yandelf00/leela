import re
import requests

url = "https://www.asciiart.eu/animals/aardvarks"
res = requests.get(url)

pattern = r"<pre.*?>(.*?)</pre>"

ascii_art_matches = re.findall(pattern, res.text, re.DOTALL)

ascii_art = [match.strip() for match in ascii_art_matches]

for art in ascii_art:
    print(art)
    print("-" * 40)  
