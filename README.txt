
CHANGES

Version 2.0.1

1. Commands are now allocated inside the ../configs/commands.json. If not found, the application stops.
2. The command now needs to specify the OS (Operating Systems) name.
2. The log path is now configurable and the default path is ../logs/
3. The log filename will use the format checklist_yyyymdhms_hostname.log. Eg: checklist_20220617151341_VM-LIT05933.log.
4. The logs will be saved inside the folder logs. Eg: ../logs/checklist_20220617151341_VM-LIT05933.log.
5. In case of success, no error will be returned.
6. The command output will be inside tags <commandOutput></commandOutput>. Eg: <commandOutput>OCSrunAchk:exit status 255</commandOutput>.
