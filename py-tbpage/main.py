import requests
import re, json


ua = 'Mozilla/5.0 (iPhone; CPU iPhone OS 11_3_1 like Mac OS X) AppleWebKit/604.1.30 (KHTML, like Gecko) Version/10.0 Safari/604.1'

_re_item_detail = re.compile(r'var _DATA_Detail = (.+);</script>')

def parseItemDetail(url):
    try:
        resp = requests.get(url, headers={'User-Agent': ua})
    except Exception as e:
        return None
    if resp.status_code is not 200:
        return None

    rs = _re_item_detail.search(resp.text)
    if not rs:
        return None

    return rs.group(1)

def main():
    tests = [
        'https://detail.tmall.com/item.htm?spm=a212k0.12153887.0.0.7187687dII24Yq&id=628244111002',
        'https://detail.tmall.com/item.htm?spm=a212k0.12153887.0.0.4d7c687dg1OiI1&id=624556842394',
    ]

    for i in tests:
        data = parseItemDetail(i)
        if not data:
            continue
        rs = json.loads(data)
        print(rs['seller'])

if __name__ == '__main__':
    main()