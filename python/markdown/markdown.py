import markdown

def parse(text):
    html=markdown.markdown(text)
    return html
