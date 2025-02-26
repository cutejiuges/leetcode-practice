class BrowserHistory:
    def __init__(self, homepage: str):
        self.urls = [homepage]
        self.currIndex = 0

    def visit(self, url: str) -> None:
        while len(self.urls) > self.currIndex + 1:
            self.urls.pop()
        self.urls.append(url)
        self.currIndex += 1

    def back(self, steps: int) -> str:
        self.currIndex = max(self.currIndex - steps, 0)
        return self.urls[self.currIndex]

    def forward(self, steps: int) -> str:
        self.currIndex = min(self.currIndex + steps, len(self.urls) - 1)
        return self.urls[self.currIndex]
