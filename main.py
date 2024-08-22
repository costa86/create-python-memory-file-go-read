import tempfile
import json
import subprocess

def create_temp_file(data):
    with tempfile.NamedTemporaryFile(mode='w',delete=False) as temp:
        temp.write(json.dumps(data))
        temp.flush()
        return temp


temp = create_temp_file({'name': 'Mary', 'age': 80, 'city':'New York'})
subprocess.run(['go',"run","main.go", temp.name])



with open(temp.name, 'r') as f:
    data = json.loads(f.read())
    print(data)