import string
import time
from datetime import datetime
from random import randint, choice

count_rows = 500000

t0 = time.time()
outfile = 'csv_file.csv'


def generate_row():
    r = '%s,%s,%s,%s\n' % (
        ''.join(choice(string.ascii_lowercase) for _ in range(10)),
        ''.join(choice(string.ascii_lowercase) for _ in range(10)),
        datetime.fromtimestamp(randint(1, int(time.time()))).strftime('%Y-%m-%d %H:%M:00'),
        ''.join(choice(string.ascii_lowercase) for _ in range(5)) + choice(['@mail.ru', '@gmail.com', '@yandex.ru'])
    )
    return r


with open(outfile, 'w') as f:
    rows = [generate_row() for row in range(count_rows)]
    f.writelines(rows)
tdelta = time.time() - t0
print(tdelta)  # 26sec for 500000
