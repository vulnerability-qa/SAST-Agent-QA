# CWE-22: Zip Slip
import zipfile, os
def extract_zip(zip_path, dest):
    with zipfile.ZipFile(zip_path) as zf:
        for member in zf.namelist():
            zf.extract(member, dest)  # ../../../etc/cron.d/backdoor
