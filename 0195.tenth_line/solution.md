```
awk 'NR==10 {print}' file.txt
```

```
cat file.txt | awk 'NR==10'
```

```
tail -n + 10 file.txt | head -n 1
```