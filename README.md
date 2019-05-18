# dynamic-page-gather

Data collecting from dynamic pages tool.

### Usage

Simple command:

```bash
dynamic-page-gather \
    --target-url=https://some.site/live-data/ \
    --gather-period=250 \
    --keep-open=10 \
    --output-file=/home/user001/newest-copy.html
```

DPG tool will render page, keep it open 10 seconds 
and every 0.25 seconds (250 millis) will save html content to file 
on disk (~/newest-copy.html)
