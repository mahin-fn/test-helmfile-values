{{- define "chart.envFrom" -}}
envFrom:
- configMapRef:
    name: {{ template "chart.fullname" . }}-configmap
{{- end -}}
